AKM: Aws access Key Manager
===

[![Build Status](https://travis-ci.org/ohsawa0515/akm.svg?branch=master)](https://travis-ci.org/ohsawa0515/akm)
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
$ akm

NAME:
   akm - A simple AWS access keys manager

USAGE:
   actions [global options] command [command options] [arguments...]

VERSION:
   0.1.0

AUTHOR:
   Shuichi Ohsawa <ohsawa0515@gmail.com>

COMMANDS:
     init, i     Initialize akm config file for the first time usage
     ls, l       List all AWS credentials profile
     use, u      Set specific AWS credential in environment values
     current, c  Show current profile name
     echo, e     Show the AWS key or region with the specified profile name
     configure   Configure AWS credentials
     clear, C    Clear the environment variable of AWS credentials
     delete, d   Delete profile from AWS credentials file
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Initialize(init)

Initialize akm config file for the first time usage. After execution, `.akm.toml` is created.


```console
$ akm init
.akm.toml is created in ~/.akm.toml
```

## List(ls)

List all AWS credentials profile.

```console
$ akm ls
default
account1
account2
```

## Use

Set specific AWS credential in environment values.

```console
# Command wapper
$ akm use <PROFILE> <Some command> [args...]

## For example
$ akm use account1 terraform apply
```

Import variables into your environment by eval.

```console
$ eval $(akm use account1)

$ env | grep AWS
AWS_ACCESS_KEY_ID=AKIAI44QH8DHBEXAMPLE
AWS_SECRET_ACCESS_KEY=je7MtGbClwBF/2Zp9Utk/h3yCo8nvbEXAMPLEKEY
AWS_DEFAULT_REGION=us-east-1
```

## Current

Show current profile name.

```console
$ akm use account1
$ akm current
account1
```

## Echo

Show the AWS key or region with the specified profile name.

```console
$ akm echo PROFILE SETTING or REGION

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

Configure AWS credentials like `aws configure --profile PROFILE_NAME`.
The set parameters are **overwritten** and saved in the credential file.

```console
$ akm configure foo
✔ AWS Access Key ID [None]: AKIAIOSFODNN7EXAMPLE
✔ AWS Secret Access Key [None]: ****************************************
✔ Region: US East (N. Virginia) [us-east-1]
? Overwrite ~/.aws/credentials? [y/N] y
```

## Clear

Clear the environment variable of AWS credentials.

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

Delete profile from AWS credentials file. When the profile is deleted, the credentials file is **overwritten**.

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

