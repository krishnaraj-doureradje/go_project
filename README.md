
# Generate Swagger Files

```bash
$ cd cmd/
$ swag init -g ./main.go -d ../cmd,../handlers,../models,../routes -o ../docs

```

## Run application 
```bash
$  go build cmd/main.go
$ ./main

```

## Linter and auto fix
```bash
$ golangci-lint run --fix
```

---

**Notes:**

- `golangci-lint` is a linter aggregator; it can auto-fix issues only if formatters like `gofmt`, `goimports`, or `gofumpt` are enabled in its config.
- `golangci-lint run` alone does not format code; use `--fix` with proper config to apply automatic formatting.
- For guaranteed formatting, explicitly run `gofmt -w .`.
- Configure `.golangci.yml` to enable formatters for auto-fix support.