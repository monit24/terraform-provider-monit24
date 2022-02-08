# Contributing

## Building

To build the binary:

```
make build
```

## Installing locally

To use the locally built provider, you have to place the binary in a special directory in `~/.terraform.d/plugins`.

See the `make install` definition for details.

```make
install: build
        mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
        mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

```

Then, in your Terraform definition, add:

```tf
terraform {
  required_providers {
    monit24 = {
      version = ">= 0.1.0"
      source  = "monit24.pl/monit24/monit24"
    }
  }
}
```

## Testing

Run unit tests with:

```
make test
```

Run acceptance tests with:

```
make testacc
```

## Documentation

Markdown [docs](./docs) are generated based on [examples](./examples) and [templates](./templates). To regenerate, run:

```
make docs
```
