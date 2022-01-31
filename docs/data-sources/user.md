---
page_title: "jumpcloud_user Data Source - terraform-provider-jumpcloud"
subcategory: ""
description: |-
  Use this data source to get information about a JumpCloud User.
---

# Data Source `jumpcloud_user`

Use this data source to get information about a JumpCloud User.

## Example Usage

```terraform
data "jumpcloud_user" "example" {
  email = "john.doe@acme.org"
}
```

```terraform
data "jumpcloud_user" "example" {
  username = "john.doe"
}
```


## Schema

### Optional

- **id** (String) he ID of this resource.
- **email** (String) The users e-mail address, which is also used for log ins. E-mail addresses have to be unique across all JumpCloud accounts, there cannot be two users with the same e-mail address. Example: `john.doe@acme.org`.
- **name** (String) The technical user name. See JumpCloud's [user naming conventions](https://support.jumpcloud.com/support/s/article/naming-convention-for-users1) for naming restrictions. Example: `john.doe`.

---
**NOTE**

one of the following must be set: id, username, email

---


