package main

import (
	"os"

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

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = "awsreset"
	app.Usage = usage
	app.Commands = []cli.Command{
		Reset,
		Grpc,
	}
}

func main() {
	app.Run(os.Args)
}
