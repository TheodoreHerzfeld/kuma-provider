---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "uptime-kuma_user Resource - terraform-provider-uptime-kuma"
subcategory: ""
description: |-
  User Resource
---

# uptime-kuma_user (Resource)

User Resource

## Example Usage

```terraform
resource "uptime-kuma_user" "my_user" {
  password = "...my_password..."
  username = "...my_username..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `password` (String) Requires replacement if changed.
- `username` (String) Requires replacement if changed.

### Read-Only

- `created_at` (String)
- `id` (Number) The ID of this resource.
- `last_visit` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import uptime-kuma_user.my_uptime-kuma_user ""
```