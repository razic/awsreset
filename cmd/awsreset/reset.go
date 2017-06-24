package main

import (
	"github.com/urfave/cli"
)

// Reset is a CLI command to reboot a collection of tagged EC2 instances
var Reset = cli.Command{
	Name:   "reset",
	Usage:  "reboots a collection of tagged EC2 instances",
	Action: reset,
}

func reset(c *cli.Context) error {
	return nil
}
