
# Generate Swagger Files

```bash
# From project root directory
$ swag init -g ./cmd/app/main.go -o ./docs

```

## Run application 
```bash
$  go build cmd/app/main.go
$ ./main

```

## Linter and auto fix
```bash
$ golangci-lint run --fix
$ go fmt ./...
$ go vet ./...
```

---

**Notes:**

- `golangci-lint` is a linter aggregator; it can auto-fix issues only if formatters like `gofmt`, `goimports`, or `gofumpt` are enabled in its config.
- `golangci-lint run` alone does not format code; use `--fix` with proper config to apply automatic formatting.
- For guaranteed formatting, explicitly run `gofmt -w .`.
- Configure `.golangci.yml` to enable formatters for auto-fix support.