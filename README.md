# Terraform Provider GPCore

This repository is a Terraform provider for [GPCore](https://g-portal.cloud).
It is intended to allow an Infrastructure-as-Code definition of the GPCore services.

Implemented Resources:
- [x] `gpcore_node` - The gpcore Node resource
- [x] `gpcore_project` - The gpcore Project resource
- [x] `gpcore_project_image` - The gpcore Project Image resource (Custom image)
- [x] `gpcore_sshkey` - The gpcore SSH-Key resource
- [x] `gpcore_billing_profile` - The gpcore Billing Profile resource

Implemented Data sources:
- [x] `gpcore_project` - The gpcore Project data source (read-only)
- [x] `gpcore_flavour` - The gpcore Flavour data source
- [x] `gpcore_datacenter` - The gpcore Datacenter data source
- [x] `gpcore_image` - The gpcore Image data source (Official images)


## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.18

## Building The GPCore Provider manually

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Using the provider

This provider is published on the [Terraform Registry](https://registry.terraform.io/providers/g-portal/gpcore/latest).

To use the gpcore Terraform Provider, all you need to do is to reference the published provider inside your terraform configuration.

```hcl
terraform {
  required_providers {
    gpcore = {
      source = "G-PORTAL/gpcore"
      # Ensure to use the latest version of the provider
      version = "0.1.4"
    }
  }
}
```

The full documentation for the provider can be found [here](https://registry.terraform.io/providers/g-portal/gpcore/latest/docs) or inside the `docs/` directory.

## Developing the Provider

If you wish to contribute to the gpcore Terraform Provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To reference the custom version of the provider built instead of the hosted one, setting `dev_override` inside the `$HOME/.terraformrc` file becomes useful:

```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/g-portal/gpcore" = "~/go/bin"
  }
  direct {}
}
```

To generate or update documentation, run `go generate`.
