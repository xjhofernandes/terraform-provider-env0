---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "env0_project Data Source - terraform-provider-env0"
subcategory: ""
description: |-
  
---

# env0_project (Data Source)



## Example Usage

```terraform
data "env0_project" "default_project" {
  name = "Default Organization Project"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **description** (String) textual description of the project
- **id** (String) id of the project
- **name** (String) the name of the project

### Read-Only

- **created_by** (String) textual description of the entity who created the project
- **role** (String) role of the authenticated user (through api key) in the project

