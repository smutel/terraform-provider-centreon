# terraform-provider-centreon

[![Build Status](https://travis-ci.org/smutel/terraform-provider-centreon.svg?branch=master)](https://travis-ci.org/smutel/terraform-provider-centreon)
[![Go Report Card](https://goreportcard.com/badge/github.com/smutel/terraform-provider-centreon)](https://goreportcard.com/report/github.com/smutel/terraform-provider-centreon)
[![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

Terraform provider to setup [Centreon Web](https://www.centreon.com/en/solutions/centreon/) application.

## Requirements

* General developper tools like make, bash, ...
* Go (to build the provider)
* Terraform (to use the provider)

## Building the provider

Clone repository to: ``$GOPATH/src/github.com/smutel/terraform-provider-centreon``

```bash
$ mkdir -p $GOPATH/src/github.com/smutel
$ cd $GOPATH/src/github.com/smutel
$ git clone git@github.com:smutel/terraform-provider-centreon.git
```

Enter the provider directory and build the provider

```bash
$ cd $GOPATH/src/github.com/smutel/terraform-provider-centreon
$ make build
```

## Installing the provider

You can install the provider manually in your global terraform provider folder 
or you can also use the makefile to install the provider in your local provider folder:

```bash
$ make localinstall
==> Creating folder terraform.d/plugins/linux_amd64
mkdir -p ~/.terraform.d/plugins/linux_amd64
==> Installing provider in this folder
cp terraform-provider-centreon ~/.terraform.d/plugins/linux_amd64
```

## Using the provider

The definition of the provider is optional.  
All the paramters could be setup by environment variables.  

```hcl
provider centreon {
  # Environment variable CENTREON_URL
  url = "http://127.0.0.1"
  
  # Environment variable CENTREON_ALLOW_UNVERIFIED_SSL
  allow_unverified_ssl = false
  
  # Environment variable CENTREON_USER
  user = "admin"
  
  # Environment variable CENTREON_PASSWORD
  password = "centreon"
}
```

For further information, check this [documentation](docs/Provider.md)

## Contributing to this project

To contribute to this project I suggest to use the [centreon style guides](https://github.com/centreon/centreon/blob/master/CONTRIBUTING.md#centreon-style-guides).  
This project is still in progress and is linked to this other project [go-centreon](https://github.com/smutel/go-centreon).

## Examples

You can find some examples in the examples folder.  
Each example can be executed directly with command terraform init & terraform apply.  
You can set different environment variables for your test:
* CENTREON_URL to define the URL (and optionally the port) | DEFAULT=http://127.0.0.1
* CENTREON_ALLOW_UNVERIFIED_SSL to avoid checking the SSL certs (true or false) | DEFAULT=false
* CENTREON_USER to define the user | DEFAULT=admin
* CENTREON_PASSWORD to define the password | DEFAULT=centreon

```bash
$ export CENTREON_URL="http://10.164.48.254:8080"
$ export CENTREON_ALLOW_UNVERIFIED_SSL="true"
$ export CENTREON_USER="admin"
$ export CENTREON_PASSWORD="centreon"
$ cd examples/commands
$ terraform init & terraform apply
```
## Known bugs which can impact this provider

* Issue [7621](https://github.com/centreon/centreon/issues/7621)
* PR [8085](https://github.com/centreon/centreon/pull/8085)
* PR [7678](https://github.com/centreon/centreon/pull/7678)
