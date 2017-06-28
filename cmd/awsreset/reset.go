package main

import (
	"bytes"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// Reset will reboot a collection of ec2 instances by tag
func Reset(svc ec2iface.EC2API, writer io.Writer, name string) error {
	// query the ec2 api, filtering by name tag
	output, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(strings.Join([]string{"*", name, "*"}, "")),
				},
			},
		},
	})

	// bubble up any errors
	if err != nil {
		return err
	}

	// verify we have a non nil output
	if output == nil {
		return nil
	}

	// initialize a slice to hold the ids for the reboot
	ids := []*string{}

	// write out the private ip addresses, and store instance ids for reboot
	for _, r := range output.Reservations {
		if r.Instances == nil {
			continue
		}

		for _, i := range r.Instances {
			var buf bytes.Buffer

			if i.PrivateIpAddress != nil {
				buf.WriteString(*i.PrivateIpAddress)
				buf.WriteRune('\n')
				buf.WriteTo(writer)
			}

			if i.InstanceId != nil {
				ids = append(ids, i.InstanceId)
			}
		}
	}

	// reboot instances
	svc.RebootInstances(&ec2.RebootInstancesInput{InstanceIds: ids})

	return nil
}
