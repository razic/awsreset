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
}

type mockEc2Client struct {
	ec2iface.EC2API
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
	}

	return output, err
}

func TestResetReturnsErrorWhenUnableToDescribeInstances(t *testing.T) {
	svc := &mockEc2Client{}
	err := Reset(svc, os.Stdout)

	if err == nil {
		t.Fatalf("%s\n", "expected error")
	}
}

func TestResetWritesIpsToWriter(t *testing.T) {
	var buf bytes.Buffer

	svc := &mockEc2Client{}
	err := Reset(svc, &buf)
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
