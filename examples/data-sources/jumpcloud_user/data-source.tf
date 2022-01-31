data "jumpcloud_user" "example_by_email" {
  email = "john.doe@acme.org"
}

data "jumpcloud_user" "example_by_username" {
  username = "john.doe"
}