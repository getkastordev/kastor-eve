# Kastor Eve Plugin

External Vercel Eve code-generation target for [Kastor](https://github.com/weirdGuy/kastor).

The plugin owns Eve-specific validation, TypeScript generation, dependency pins, model gateway mapping, MCP connections, approval gates, runtime-tool scaffolds, and the Eve starter returned by `kastor new`. Kastor core sends canonical protocol-v1 IR and remains responsible for deterministic disk writes.

> Status: pre-release. The protocol-v1 implementation is available for integration testing but no stable binary has been released yet.

## Development

```sh
go test ./...
go vet ./...
go build ./cmd/kastor-eve
```

Point a local Kastor checkout at the binary with:

```sh
export KASTOR_PLUGIN_EVE=/absolute/path/to/kastor-eve
```

The source address is `github.com/getkastordev/kastor-eve`.

Create and lock an Eve starter with:

```sh
kastor new --from github.com/getkastordev/kastor-eve demo
```

## License

Apache-2.0
