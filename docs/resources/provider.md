---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "env0_provider Resource - terraform-provider-env0"
subcategory: ""
description: |-
  
---

# env0_provider (Resource)



## Example Usage

```terraform
resource "env0_provider" "example" {
  type        = "aws-key-example"
  description = "description example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `type` (String) type of the provider (Match pattern: ^[0-9a-zA-Z](?:[0-9a-zA-Z-]{0,30}[0-9a-zA-Z])?$). Your provider’s type is essentially it’s name, and should match your provider’s files. For example, if your binaries look like terraform-provider-aws_1.1.1_linux_amd64.zip, than your provider’s type should be aws.

### Optional

- `description` (String) description of the provider

### Read-Only

- `id` (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import env0_provider.by_id ddda7b30-6789-4d24-937c-22322754934e
terraform import env0_provider.by_type aws
```