AKM: Aws access Key Manager
===

![Test Status](https://github.com/ohsawa0515/akm/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ohsawa0515/akm)](https://goreportcard.com/report/github.com/ohsawa0515/akm)

AKM is a simple AWS access Keys Manager.

It helps you to switch multiple AWS credentials easily. And it is useful when could not specify profile of AWS credentials. e.g. Terraform, Packer, Bash scripts, PowerShell scirpts and so on.

# Installation

## Download binary

Download it from [releases page](https://github.com/ohsawa0515/akm/releases) and extract it to /usr/bin or your PATH directory.

## Using Go

If you have not installed [dep](https://github.com/golang/dep) yet, please install it.

```console
$ go get -u github.com/golang/dep/cmd/dep 
```

```console
$ go get -u github.com/ohsawa0515/akm
$ dep ensure
```


# Set AWS credentials

- `$HOME/.aws/credentials`

```ini
# For example
[default]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

[account1]
aws_access_key_id = AKIAI44QH8DHBEXAMPLE
aws_secret_access_key = je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
region = us-east-1 # optional

[account2]
aws_access_key_id = AKIAI44QH7DHBEXAMPLE
aws_secret_access_key = je7MtGbClwBF/3Zp9Utk/h4yCo8nvbEXAMPLEKEY
region = us-east-2 # optional
```

- `$HOME/.aws/config` (Optional)

```ini
[default]
region=us-west-2
output=json

[profile acountA]
region=us-east-1
output=text

[profile acountB]
region=us-east-2
output=text
```

# Usage

```console
$ akm -help

Name:
  akm - A simple AWS access keys manager

Usage:
  akm COMMAND

Version:
  0.0.1

Author:
  Shuichi Ohsawa <ohsawa0515@gmail.com>

Commands:
  clear       Delete the environment variable of AWS credentials.
  configure   Configure AWS credentials
  current     Show current profile name.
  delete      Delete profile from AWS credentials file.
  echo        Show the AWS key or region with the specified profile name.
  help        Help about any command
  init        Initialize for akm command
  list        List all AWS credentials profile.
  use         Use specific AWS credential key

Flags:
  -h, --help      Print help
  -v, --version   Print the version
```

## Initialize(init)

```console
$ akm init --help

Name:
  akm init - Initialize for akm command

Usage:
  akm init

Aliases:
  init, i

Description:
  Initialize akm command for the first time usage.
  After execution, "$HOME/.akm.toml" is created.

Global Flags:
  -h, --help   Print help
```

For example:

```console
$ akm init
.akm.toml is created in ~/.akm.toml
```

## List(ls)

```console
$ akm ls --help

Name:
  akm list - List all AWS credentials profile.

Usage:
  akm list

Aliases:
  list, ls, l

Global Flags:
  -h, --help   Print help
```

For example:

```console
$ akm ls
default
account1
account2
```

## Use

```console
$ akm use --help

Name:
  akm use - Use specific AWS credential key

Usage:
  akm use PROFILE [[ANY COMMAND]...]

Aliases:
  use, u

Description:
  Set specific AWS credential in environment values.
    - AWS_ACCESS_KEY_ID
    - AWS_SECRET_ACCESS_KEY
    - AWS_SESSION_TOKEN  (if session token is set)
    - AWS_DEFAULT_REGION (if region is set)
  If an arbitrary command was specified as an argument, store the AWS credentials key in the environment variable and then execute the command.

Examples:
  case 1) Set specific AWS credential in environment values.
    $ akm use foo
    export AWS_ACCESS_KEY_ID='xxxxxxxx';export AWS_SECRET_ACCESS_KEY='xxxxxxxxx';export AWS_DEFAULT_REGION=us-east-1

  case 2) Import variables into your environment by eval.
    $ eval $(akm use foo)
    $ env | grep AWS
    AWS_ACCESS_KEY_ID=xxxxxxxx
    AWS_SECRET_ACCESS_KEY=xxxxxxxxx
    AWS_DEFAULT_REGION=us-east-1

  case 3) Store the AWS credentials key in the environment variable and then execute the command.
    $ akm use foo terraform plan

Global Flags:
  -h, --help   Print help
```

For example:

```console
# Import variables into your environment by eval.
$ eval $(akm use account1)

$ env | grep AWS
AWS_ACCESS_KEY_ID=AKIAI44QH8DHBEXAMPLE
AWS_SECRET_ACCESS_KEY=je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
AWS_DEFAULT_REGION=us-east-1
```

```console
# Command wapper
$ akm use account1 terraform apply
```

## Current

```console
$ akm current --help

Name:
  akm current - Show current profile name.

Usage:
  akm current

Aliases:
  current, c

Global Flags:
  -h, --help   Print help
```

For example:

```console
$ akm use account1
$ akm current
account1
```

## Echo

```console
$ akm echo --help

Name:
  akm echo - Show the AWS key or region with the specified profile name.

Usage:
  akm echo PROFILE aws_access_key_id | aws_secret_access_key | region

Aliases:
  echo, e

Global Flags:
  -h, --help   Print help
```

For example:

```console
$ akm echo account1 aws_access_key_id
AKIAI44QH8DHBEXAMPLE

$ akm echo account1 aws_secret_access_key
je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY

$ akm echo account1 region
us-east-1

# Show current AWS key
$ aws use account2
$ akm echo $(akm current) aws_access_key_id
AKIAI44QH8DHBEXAMPLE2  # account2's AWS access key id
```

## Configure

```console
$ akm configure --help

Name:
  akm configure - Configure AWS credentials

Usage:
  akm configure PROFILE

Description:
  Configure AWS credentials like "aws configure --profile PROFILE".
  The set parameters are **overwritten** and saved in the credential file.

Examples:
  akm configure foo

Global Flags:
  -h, --help   Print help
```

For example:

```console
$ akm configure foo
✔ AWS Access Key ID [None]: AKIAIOSFODNN7EXAMPLE
✔ AWS Secret Access Key [None]: ****************************************
✔ Region: US East (N. Virginia) [us-east-1]
? Overwrite ~/.aws/credentials? [y/N] y
```

## Clear

```console
$ akm clear --help

Name:
  akm clear - Delete the environment variable of AWS credentials.

Usage:
  akm clear

Aliases:
  clear, C

Examples:
  $ akm clear
  unset AWS_ACCESS_KEY_ID;unset AWS_SECRET_ACCESS_KEY;unset AWS_DEFAULT_REGION;

  Delete environment variable with eval.
  $ env | grep AWS
  AWS_ACCESS_KEY_ID=xxxxxxx
  AWS_SECRET_ACCESS_KEY=xxxxxxx
  AWS_DEFAULT_REGION=us-east-1

  $ eval $(akm clear)

  $ env | grep AWS
  # empty

Global Flags:
  -h, --help   Print help
```

For example

```console
$ env | grep AWS
AWS_ACCESS_KEY_ID=AKIAI44QH8DHBEXAMPLE
AWS_SECRET_ACCESS_KEY=je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
AWS_DEFAULT_REGION=us-east-1

$ eval $(akm clear)

$ env | grep AWS
# empty
```

## Delete

```console
$ akm delete --help

Name:
  akm delete - Delete profile from AWS credentials file.

Usage:
  akm delete PROFILE

Aliases:
  delete, del, d

Description:
  Delete profile from AWS credentials file.
  When the profile is deleted, the credentials file is **overwritten**.

Examples:
  akm delete foo

Global Flags:
  -h, --help   Print help
```

For example:

```console
$ akm delete foo
? Remove profile: foo, overwrite ~/.aws/credentials? [y/N] y
```

# Inspired by

- [https://github.com/TimothyYe/skm](https://github.com/TimothyYe/skm)
- [https://github.com/fujiwara/aswrap](https://github.com/fujiwara/aswrap)

# Contribution

1. Fork ([https://github.com/ohsawa0515/akm/fork](https://github.com/ohsawa0515/akm/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request

# License

See [LICENSE](https://github.com/ohsawa0515/akm/blob/master/LICENSE).

# Author

Shuichi Ohsawa (@ohsawa0515)

