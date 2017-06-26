package main

import (
	"errors"
	"runtime"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

var cases = map[string][]interface{}{
	"command-line-arguments.TestResetReturnsErrorWhenUnableToDescribeInstances": []interface{}{
		nil, errors.New("whoa"),
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

	// look up arguments for this test case by caller function name
	fpcs := make([]uintptr, 1)
	runtime.Callers(3, fpcs)
	fun := runtime.FuncForPC(fpcs[0] - 1)
	c := cases[fun.Name()]
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
	err := Reset(svc)

	if err == nil {
		t.Fatalf("%s\n", "expected error")
	}
}
