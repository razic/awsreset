package main

import (
	"bytes"
	"errors"
	"os"
	"runtime"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

var cases = map[string][]interface{}{
	"command-line-arguments.TestResetReturnsErrorWhenUnableToDescribeInstances": []interface{}{
		nil,
		errors.New(""),
	},
	"command-line-arguments.TestResetWritesIpsToWriter": []interface{}{
		&ec2.DescribeInstancesOutput{
			Reservations: []*ec2.Reservation{
				{
					Instances: []*ec2.Instance{
						{PrivateIpAddress: aws.String("1")},
						{PrivateIpAddress: aws.String("2")},
						{PrivateIpAddress: aws.String("3")},
					},
				},
			},
		},
		nil,
	},
	"command-line-arguments.TestResetFiltersOnNameTag": []interface{}{
		nil,
		nil,
	},
	"command-line-arguments.TestResetRebootsInstances": []interface{}{
		&ec2.DescribeInstancesOutput{
			Reservations: []*ec2.Reservation{
				{
					Instances: []*ec2.Instance{
						{InstanceId: aws.String("1")},
						{InstanceId: aws.String("2")},
						{InstanceId: aws.String("3")},
					},
				},
			},
		},
		nil,
	},
	"command-line-arguments.TestResetDoesntRebootsInstancesDuringDryRun": []interface{}{
		nil,
		nil,
	},
}

type mockEc2Client struct {
	ec2iface.EC2API
}

func (m *mockEc2Client) RebootInstances(input *ec2.RebootInstancesInput) (*ec2.RebootInstancesOutput, error) {
	var (
		output *ec2.RebootInstancesOutput
		err    error
	)

	fpcs := make([]uintptr, 1)
	runtime.Callers(3, fpcs)
	fun := runtime.FuncForPC(fpcs[0] - 1)
	c := cases[fun.Name()]

	// look up return arguments for this test case by caller function name
	if c != nil {
		cases[fun.Name()] = append(c, input)
	}

	return output, err
}

func (m *mockEc2Client) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	var (
		output *ec2.DescribeInstancesOutput
		err    error
	)

	fpcs := make([]uintptr, 1)
	runtime.Callers(3, fpcs)
	fun := runtime.FuncForPC(fpcs[0] - 1)
	c := cases[fun.Name()]

	// look up return arguments for this test case by caller function name
	if c != nil {
		if c[0] != nil {
			output = c[0].(*ec2.DescribeInstancesOutput)
		}

		if c[1] != nil {
			err = c[1].(error)
		}

		cases[fun.Name()] = append(c, input)
	}

	return output, err
}

func TestResetReturnsErrorWhenUnableToDescribeInstances(t *testing.T) {
	svc := &mockEc2Client{}
	err := Reset(svc, os.Stdout, "", true)

	if err == nil {
		t.Fatalf("%s\n", "expected error")
	}
}

func TestResetWritesIpsToWriter(t *testing.T) {
	var buf bytes.Buffer

	svc := &mockEc2Client{}
	err := Reset(svc, &buf, "", true)
	lines := bytes.Split(buf.Bytes(), []byte{'\n'})

	if err != nil {
		t.Fatalf("%s\n", "unexpected error")
	}

	if string(lines[0]) != "1" {
		t.Fatalf("%s\n", "unexpected output")
	}

	if string(lines[1]) != "2" {
		t.Fatalf("%s\n", "unexpected output")
	}

	if string(lines[2]) != "3" {
		t.Fatalf("%s\n", "unexpected output")
	}
}

func TestResetFiltersOnNameTag(t *testing.T) {
	var input *ec2.DescribeInstancesInput

	name := "etcd"
	svc := &mockEc2Client{}
	err := Reset(svc, os.Stdout, name, true)

	if err != nil {
		t.Fatalf("%s\n", "unexpected error")
	}

	c := cases["command-line-arguments.TestResetFiltersOnNameTag"]

	if c == nil || len(c) < 3 {
		t.Fatalf("%s\n", "couldnt find input for test case")
	}

	input = c[2].(*ec2.DescribeInstancesInput)

	if len(input.Filters) != 1 {
		t.Fatalf("expected only 1 filter, got %d\n", len(input.Filters))
	}

	filter := input.Filters[0]

	if *filter.Name != "tag:Name" {
		t.Fatalf("expected name to be tag:Name, got %s\n", filter.Name)
	}

	if len(filter.Values) != 1 {
		t.Fatalf("expected a single value in the filter\n")
	}

	val := filter.Values[0]

	if *val != "*etcd*" {
		t.Fatalf("%s != %s\n", *val, "*etcd*")
	}
}

func TestResetRebootsInstances(t *testing.T) {
	var input *ec2.RebootInstancesInput

	svc := &mockEc2Client{}
	err := Reset(svc, os.Stdout, "", false)

	if err != nil {
		t.Fatalf("%s\n", "unexpected error")
	}

	c := cases["command-line-arguments.TestResetRebootsInstances"]

	if c == nil {
		t.Fatalf("%s\n", "unable to find test case params")
	}

	if len(c) < 4 {
		t.Fatalf("%s\n", "unable to find input for test case")
	}

	input = c[3].(*ec2.RebootInstancesInput)
	ids := input.InstanceIds

	if len(input.InstanceIds) != 3 {
		t.Fatalf("%s\n", "incorrect number of instace ids")
	}

	if *ids[0] != "1" {
		t.Fatalf("%s\n", "unexpected instance id")
	}

	if *ids[1] != "2" {
		t.Fatalf("%s\n", "unexpected instance id")
	}

	if *ids[2] != "3" {
		t.Fatalf("%s\n", "unexpected instance id")
	}
}
func TestResetDoesntRebootsInstancesDuringDryRun(t *testing.T) {
	svc := &mockEc2Client{}
	err := Reset(svc, os.Stdout, "", true)

	if err != nil {
		t.Fatalf("%s\n", "unexpected error")
	}

	c := cases["command-line-arguments.TestResetDoesntRebootsInstancesDuringDryRun"]

	if c == nil {
		t.Fatalf("%s\n", "unable to find test case params")
	}

	if len(c) >= 4 {
		t.Fatalf("%s\n", "shouldnt have found any input for test case")
	}
}
