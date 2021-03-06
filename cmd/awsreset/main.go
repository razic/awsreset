package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/urfave/cli"
)

const usage = `
                                        _
   __ ___      _____ _ __ ___  ___  ___| |_
  / _  \ \ /\ / / __| '__/ _ \/ __|/ _ \ __|
 | (_| |\ V  V /\__ \ | |  __/\__ \  __/ |_
  \__,_| \_/\_/ |___/_|  \___||___/\___|\__|

reboots instances by tag
`

func main() {
	app := cli.NewApp()

	app.Name = "awsreset"
	app.Usage = usage
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dry-run",
			Usage: "do not actually perform any reboots",
		},
		cli.StringFlag{
			Name:  "region",
			Value: "us-west-2",
			Usage: "ec2 region",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "value for tag:Name filter",
		},
	}
	app.Action = func(c *cli.Context) error {
		sess := session.Must(session.NewSession(aws.NewConfig().WithRegion(c.GlobalString("region"))))
		svc := ec2.New(sess)
		err := Reset(svc, os.Stdout, c.GlobalString("name"), c.GlobalBool("dry-run"))

		if err != nil {
			return cli.NewExitError(err.Error(), 1)
		}

		return nil
	}

	app.Run(os.Args)
}
