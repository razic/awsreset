package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/urfave/cli"
)

// Reset is a CLI command to reboot a collection of tagged EC2 instances
var Reset = cli.Command{
	Name:   "reset",
	Usage:  "reboots a collection of tagged EC2 instances",
	Action: reset,
}

func reset(c *cli.Context) error {
	svc := ec2.New(Session)
	output, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})

	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v\n", output)

	return nil
}
