---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "env0_aws_credentials Resource - terraform-provider-env0"
subcategory: ""
description: |-
  
---

# env0_aws_credentials (Resource)



## Example Usage

```terraform
resource "env0_aws_credentials" "credentials" {
  name        = "example"
  arn         = "Example role ARN"
  external_id = "Example external id"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **arn** (String) the aws role arn
- **external_id** (String, Sensitive) the aws role external id
- **name** (String) name for the credentials

### Optional

- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import env0_aws_credentials.by_id d31a6b30-5f69-4d24-937c-22322754934e
terraform import env0_aws_credentials.by_name ProductionCredentials
```