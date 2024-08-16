# terraform-provider-lambdalabs


Create a local provider configuration file:
Create a file named .terraformrc in your home directory (or terraform.rc in Windows) with the following content:

```
provider_installation {
  dev_overrides {
    "WillBeebe/lambdalabs" = "/Users/wbeebe/repos/_current/terraform-provider-lambdalabs/bin"
  }
  direct {}
}
```

`terraform init -upgrade`


https://github.com/hashicorp/terraform-provider-scaffolding/blob/main/internal/provider/resource_scaffolding.go
