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