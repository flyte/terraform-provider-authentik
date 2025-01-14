<p align="center">
    <img src="https://goauthentik.io/img/icon_top_brand_colour.svg" height="150" alt="authentik logo">
</p>

---

[![](https://img.shields.io/discord/809154715984199690?label=Discord&style=for-the-badge)](https://discord.gg/jg33eMhnj6)
[![Code Coverage](https://img.shields.io/codecov/c/gh/goauthentik/terraform-provider-authentik?style=for-the-badge)](https://codecov.io/gh/goauthentik/terraform-provider-authentik)
[![Latest version](https://img.shields.io/github/v/tag/goauthentik/terraform-provider-authentik?style=for-the-badge)](https://registry.terraform.io/providers/goauthentik/authentik/latest)
[![CI Build status](https://img.shields.io/github/actions/workflow/status/goauthentik/terraform-provider-authentik/test.yml?branch=main&style=for-the-badge)](https://github.com/goauthentik/terraform-provider-authentik/actions)

# Terraform Provider authentik

Tested against authentik main and stable, on terraform 1.2.1

Run the following command to build the provider

```shell
make build
```

### Generate Documentation

Run `make` from the project root to regenerate the latest provider documentation

## Testing

In `./tests`, run `docker compose up -d` to start a testing instance with a fixed token.

Source `./tests/.env` and run go tests to start testing against the local instance.

## Versioning

This provider's version is based on the authentik version it's tested against.

Provider version 2021.8.1 is tested against 2021.8.x.

Provider version 2021.8.2 is tested against 2021.8.x but has some bugfixes.
