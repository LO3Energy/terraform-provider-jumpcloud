---
page_title: "jumpcloud_user_group_association Resource - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Provides a resource for associating a JumpCloud user group to objects like SSO applications, G Suite, Office 365, LDAP and more.
---

# Resource `jumpcloud_user_group_association`

Provides a resource for associating a JumpCloud user group to objects like SSO applications, G Suite, Office 365, LDAP and more.



## Schema

### Required

- **group_id** (String) The ID of the `resource_user_group` resource.
- **object_id** (String) The ID of the object to associate to the group.
- **type** (String) The type of the object to associate to the given group. Possible values: `active_directory`, `application`, `command`, `g_suite`, `ldap_server`, `office_365`, `policy`, `radius_server`, `system`, `system_group`.

### Optional

- **id** (String) The ID of this resource.


