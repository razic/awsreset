package main

import (
	"bytes"
	"io"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// Reset will reboot a collection of ec2 instances by tag
func Reset(svc ec2iface.EC2API, writer io.Writer) error {
	output, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})

	if err != nil {
		return err
	}

	for _, r := range output.Reservations {
		for _, i := range r.Instances {
			var buf bytes.Buffer

			buf.WriteString(*i.PrivateIpAddress)
			buf.WriteRune('\n')
			buf.WriteTo(writer)
		}
	}

	return nil
}
