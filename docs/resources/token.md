---
page_title: "authentik_token Resource - terraform-provider-authentik"
subcategory: "Identity"
description: |-
---

# authentik_token (Resource)

## Example Usage

```terraform
# Create/manage API/App Password tokens

resource "authentik_token" "default" {
  identifier = "my-token"
  user = authentik_user.some_user.id
  intent = "api"
  # If this is not set then the actual token won't be retrieved
  retrieve_key = true
}

# Use authentik_token.default.key
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **identifier** (String)
- **user** (Number)

### Optional

- **description** (String)
- **expires** (String)
- **expiring** (Boolean) Defaults to `true`.
- **id** (String) The ID of this resource.
- **intent** (String) Defaults to `api`.
- **retrieve_key** (Boolean) Defaults to `false`.

### Read-Only

- **expires_in** (Number)
- **key** (String, Sensitive)

