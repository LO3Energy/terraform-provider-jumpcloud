---
page_title: "jumpcloud_user_group_membership Resource - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Provides a resource for managing user group memberships.
---

# Resource `jumpcloud_user_group_membership`

Provides a resource for managing user group memberships.

## Example Usage

```terraform
data "jumpcloud_user_group" "example" {
  name = "My User Group"
}

data "jumpcloud_user" "john_doe" {
  email      = "john.doe@acme.org"
}


resource "jumpcloud_user_group_membership" "example" {
  user_id  = data.jumpcloud_user.john_doe.id
  group_id = data.jumpcloud_user_group.example.id
}
```

## Schema

### Required

- **group_id** (String) The ID of the `resource_user_group` object.
- **user_id** (String) The ID of the `resource_user` object.

### Optional

- **id** (String) The ID of this resource.


