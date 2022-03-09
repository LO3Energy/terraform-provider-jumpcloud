---
page_title: "jumpcloud_user_group Data Source - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Use this data source to get information about a JumpCloud Group.
---

# Data Source `jumpcloud_user_group`

Use this data source to get information about a JumpCloud Group.

## Example Usage

```terraform
data "jumpcloud_user_group" "example" {
  name = "My User Group"
}
```


## Schema

### Optional

- **id** (String) he ID of this resource.
- **name** (String) The technical group name.

---
**NOTE**

one of the following must be set: id, name

---


