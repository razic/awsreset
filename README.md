# awsreset

`awsreset` is a simple tool that can be used to reset a collection of tagged
EC2 instances.

# Installation

[Download](https://github.com/razic/awsreset/releases) the latest release (pre-built binary) for your platform, and place
it in your `$PATH`.

### Supported Platforms

* Mac (amd64)
* Linux (amd64)

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
   1.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --dry-run       do not actually perform any reboots
   --region value  ec2 region (default: "us-west-2")
   --name value    value for tag:Name filter
   --help, -h      show help
   --version, -v   print the version
```

### Authentication to AWS

Authentication to AWS is supported through the use of the official Amazon Go
SDK. The SDK provides many methods for authentication. If you are unfamiliar
with traditional authentication methods with AWS, please read
https://github.com/aws/aws-sdk-go#configuring-credentials.

# Development

If you are interested in contributing to `awsreset`, you will need all of the
Go dependencies in your `$GOPATH`. For the purposes of this project, I chose
not to use a dependency manager.

Running `make` will create the following executables:

* `./bin/awsreset-darwin-amd64-v1.0.0` (Mac)
* `./bin/aws-reset-linux-amd64-v1.0.0` (Linux)

## Tests

`awsreset` uses the standard Go `testing` library. Execute `make test` to run
the tests.
