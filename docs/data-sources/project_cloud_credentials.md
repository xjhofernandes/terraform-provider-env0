---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "env0_project_cloud_credentials Data Source - terraform-provider-env0"
subcategory: ""
description: |-
  
---

# env0_project_cloud_credentials (Data Source)



## Example Usage

```terraform
data "env0_project_cloud_credentials" "project_cloud_credentials" {
  project_id = "pid1"
}

output "pid1_credential_0_id" {
  value = data.env0_project_cloud_credentials.project_cloud_credentials.ids.0
}

output "pid1_credential_1_id" {
  value = data.env0_project_cloud_credentials.project_cloud_credentials.ids.1
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `project_id` (String) the project id for listing the cloud credentials

### Read-Only

- `id` (String) The ID of this resource.
- `ids` (List of String) a list of cloud credentials (ids) associated with the project

