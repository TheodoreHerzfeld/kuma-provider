# Kuma-provider

This project aims to create a terraform provider (named "uptime-kuma" in terraform code) for the [Uptime-Kuma selfhosted monitoring system](https://github.com/louislam/uptime-kuma).

This repository is built on the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). The template repository built on the [Terraform Plugin SDK](https://github.com/hashicorp/terraform-plugin-sdk) can be found at [terraform-provider-scaffolding](https://github.com/hashicorp/terraform-provider-scaffolding) and is used as the basis for this project.

This project has not yet been published to the Terraform Registry. See instructions below for local development/usage.

### NOTE
This project's current goal is not to get a working provider. Without the official API, and the webapi layer in use being over a year out of development
and showing some bugs, it is unlikely to have a working provider. For now, the goal is to create the framework for when the official API is released, with
focuses on creating schemas and resources setup for later use.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.22
- [Docker](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/)


## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Adding Dependencies

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```shell
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.

## Using the provider

See [Prepare Terraform for local provider install](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install) for information on how to point your Terraform binary to your local build
for development (using the `~/.terraformrc` file).

While not published yet, the registry url will probably be `hashicorp.com/theodoreherzfeld/uptime-kuma`.

## Developing the Provider

Start by setting up the `docker-compose.yml` file:

1. set `KUMA_USERNAME` and `KUMA_PASSWORD` to what you INTEND to use as the login credentials for your development uptime-kuma instance
2. run `docker compose up` to start the development environment. Connect to the development uptime-kuma instance at `localhost:3069` and 
    setup the admin account
3. connect to the webapi at `localhost:8000` for information about the API during development
4. ?
5. profit!

Until the official API is released, the provider configuration should point to the API, not directly to the uptime-kuma instance. The credentials
for access are set by the environment variables attached to the API service in `docker-compose.yml`.

---

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

```shell
make testacc
```

## Todos:
1. github actions, publish to registry (probably gonna wait a while on this)

## Features so far:

Provider:
* password-based login

Data sources:
1. user

Resources:
1. monitor - broken as hell - suspect issues with the API laye

<!-- No SDK Installation -->
<!-- No SDK Example Usage -->
<!-- No SDK Available Operations -->
<!-- Start Summary [summary] -->
## Summary


<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents

* [Installation](#installation)
* [Available Resources and Data Sources](#available-resources-and-data-sources)
* [Testing the provider locally](#testing-the-provider-locally)
<!-- End Table of Contents [toc] -->

<!-- Start Installation [installation] -->
## Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    uptime-kuma = {
      source  = "TheodoreHerzfeld/uptime-kuma"
      version = "0.1.0"
    }
  }
}

provider "uptime-kuma" {
  # Configuration options
}
```
<!-- End Installation [installation] -->

<!-- Start Available Resources and Data Sources [operations] -->
## Available Resources and Data Sources

### Resources

* [uptime-kuma_monitor](docs/resources/monitor.md)
* [uptime-kuma_user](docs/resources/user.md)
### Data Sources

* [uptime-kuma_monitor](docs/data-sources/monitor.md)
* [uptime-kuma_monitors](docs/data-sources/monitors.md)
* [uptime-kuma_user](docs/data-sources/user.md)
<!-- End Available Resources and Data Sources [operations] -->

<!-- Start Testing the provider locally [usage] -->
## Testing the provider locally

#### Local Provider

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```

#### Compiled Provider

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

1. Execute `go build` to construct a binary called `terraform-provider-uptime-kuma`
2. Ensure that the `.terraformrc` file is configured with a `dev_overrides` section such that your local copy of terraform can see the provider binary

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/TheodoreHerzfeld/uptime-kuma" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```
<!-- End Testing the provider locally [usage] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

r