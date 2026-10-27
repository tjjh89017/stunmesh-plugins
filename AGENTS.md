# AGENTS.md

## Purpose

Reference storage plugins for stunmesh-go, run as external processes over
the exec (JSON) and shell (shell variables) plugin protocols. stunmesh-go's
built-in plugins cover the same backends. There are no releases; users build
from source.

## Layout

- `cloudflare/`: exec plugin in Go, its own module
  `github.com/tjjh89017/stunmesh-plugins/cloudflare`, with `go test`.
- `cloudflare-shell/`, `opendht/`, `opendht-shell/`: POSIX sh plugins, each
  with `smoke_test.sh` against a fake backend.
- `Makefile`: runs a target in every top-level directory that has a Makefile.
- Each plugin Makefile uses `PLUGIN` (not `APP`) for the `stunmesh-<name>`
  binary and has `build`, `clean`, `test`, `install`, `uninstall`.

## Commands

```bash
make build                 # build all plugins in place
make test                  # go test + every smoke_test.sh
make -C <dir> test         # one plugin
make clean
```

Smoke tests need `curl`, `python3` and `jq`. They run offline.

## Relation to stunmesh-go

- The protocols are defined by stunmesh-go and documented at
  https://docs.stunmesh.dev/plugins/exec-protocol and
  https://docs.stunmesh.dev/plugins/shell-protocol.
- `cloudflare/` imports `github.com/tjjh89017/stunmesh-go/pluginapi` at a
  published version. Dependabot opens a PR when stunmesh-go releases; to
  bump by hand, run `go get github.com/tjjh89017/stunmesh-go@vX.Y.Z && go
  mod tidy` in `cloudflare/`.
- To test against a local stunmesh-go checkout, uncomment the `replace`
  line in `cloudflare/go.mod`. Do not commit it uncommented.

## Commits

- Sign off every commit: `git commit -s`.
- Use conventional commits: `type: description`, with type `feat`, `fix`,
  `docs`, `refactor`, `test`, `chore` or `ci`. Keep the first line under 72
  characters.
- Do not add AI attribution or AI co-author trailers.
- Stage files by explicit path.

## License

Apache-2.0 for everything in this repository (see `LICENSE`). Contributors
sign the stunmesh CLA; the `CLA` workflow checks it on pull requests.
