package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// Reset will reboot a collection of ec2 instances by tag
func Reset(svc ec2iface.EC2API) error {
	output, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})

	if err != nil {
		return err
	}

	fmt.Printf("%v\n", output)

	return nil
}
