# awsreset

`awsreset` is a simple tool that can be used to reset a collection of tagged
EC2 instances.

# Usage

```bash
NAME:
   awsreset -
                                        _
   __ ___      _____ _ __ ___  ___  ___| |_
  / _  \ \ /\ / / __| '__/ _ \/ __|/ _ \ __|
 | (_| |\ V  V /\__ \ | |  __/\__ \  __/ |_
  \__,_| \_/\_/ |___/_|  \___||___/\___|\__|

reboots instances by tag


USAGE:
   awsreset [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dry-run       do not actually perform any reboots
   --region value  ec2 region (default: "us-west-2")
   --help, -h      show help
   --version, -v   print the version
```

# Development

If you are interested in contributing to `awsreset` all of the Go dependencies
in your `$GOPATH`. For the purposes of this project, I chose not to use a
dependency manager.

Running `make` will create the following executables:

* `./bin/awsreset-darwin-amd64` (Mac)
* `./bin/aws-reset-linux-amd64` (Linux)

## Tests

`awsreset` uses the standard Go `testing` library. Execute `make test` to run
the tests.
