package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/urfave/cli"
)

const usage = `
                                        _
   __ ___      _____ _ __ ___  ___  ___| |_
  / _  \ \ /\ / / __| '__/ _ \/ __|/ _ \ __|
 | (_| |\ V  V /\__ \ | |  __/\__ \  __/ |_
  \__,_| \_/\_/ |___/_|  \___||___/\___|\__|

tool to reset aws instances
`

// Session is the global session object for the AWS api
var Session *session.Session

func main() {
	app := cli.NewApp()
	app.Name = "awsreset"
	app.Usage = usage
	app.Commands = []cli.Command{
		Reset,
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "region",
		},
	}
	app.Before = func(c *cli.Context) error {
		Session = session.Must(session.NewSession(aws.NewConfig().WithRegion(c.GlobalString("region"))))
		return nil
	}
	app.Run(os.Args)
}
