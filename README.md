# Monit24.pl Terraform Provider

## Usage

```tf

```

See [examples](./examples) for in-depth usage.

## Authentication

Fill the `user` and `password` directly in the provider (preferably via Terraform's variables).

```tf
provider "monit24" {
    user = "username"
    password = "topsecret"
}
```

Alternatively, provide `MONIT24_USER` and `MONIT24_PASSWORD` environment variables.

```tf
provider "monit24" {}
```

```bash
export MONIT24_USER=
export MONIT24_PASSWORD=
```

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md)