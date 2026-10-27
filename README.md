# stunmesh-plugins

Storage plugins for [stunmesh-go](https://github.com/tjjh89017/stunmesh-go)
that run as separate processes, using stunmesh-go's **exec** (JSON) and
**shell** (shell variables) plugin protocols.

stunmesh-go's built-in plugins cover the same backends (Cloudflare DNS and
OpenDHT) and need no extra binary. The plugins here are reference
implementations: use them as a starting point for your own backend, or when
you want a plugin outside the stunmesh-go binary. There are no prebuilt
releases; build from source.

## Plugins

| Directory | Protocol | Language | Description |
|-----------|----------|----------|-------------|
| [`cloudflare/`](cloudflare/) | exec | Go | Stores endpoints in Cloudflare DNS TXT records. Static binary. |
| [`cloudflare-shell/`](cloudflare-shell/) | shell | sh | Same Cloudflare storage in a shell script; needs curl, grep, sed. |
| [`opendht/`](opendht/) | exec | sh | Stores endpoints in OpenDHT through a DHT proxy REST API; needs curl, jq. |
| [`opendht-shell/`](opendht-shell/) | shell | sh | Same OpenDHT storage without jq; needs curl, base64. |

Each directory has a README with setup and configuration examples.

## Build and install

Requires Go (version from [`cloudflare/go.mod`](cloudflare/go.mod)) for the
cloudflare plugin, and `make`. The shell plugins need no build step; `make
build` copies each script to its `stunmesh-*` name and makes it executable.

```bash
make build                              # build every plugin in place
make install                            # install to /usr/local/bin
make install PREFIX=/opt/stunmesh       # or BINDIR=/some/bin
make uninstall
make clean
make list                               # show detected plugins
```

The top-level Makefile runs the same target in every directory that has a
Makefile. To work on one plugin, run `make -C <dir> <target>`. The cloudflare
plugin cross-compiles with `GOOS`/`GOARCH`, for example
`make -C cloudflare build GOOS=linux GOARCH=mipsle`.

## Tests

```bash
make test
```

This runs `go test ./...` for the cloudflare plugin and `smoke_test.sh` for
each shell plugin. The smoke tests drive the real scripts against fake
backends, so they need no network or credentials. They need `curl`,
`python3` (for the mock servers) and, for `opendht/`, `jq`. Run one test
with `make -C <dir> test`. CI runs `make test` on every push and pull
request.

## Writing your own plugin

The protocols are documented at docs.stunmesh.dev:

- [Exec plugin protocol](https://docs.stunmesh.dev/plugins/exec-protocol)
- [Shell plugin protocol](https://docs.stunmesh.dev/plugins/shell-protocol)

Use [`cloudflare/`](cloudflare/) as the reference for an exec plugin in Go,
and [`opendht-shell/`](opendht-shell/) for a shell plugin.

Notes from the existing plugins:

- **Exec exit code.** stunmesh-go treats a call as failed when the process
  exits non-zero *or* the response has `success: false`. `cloudflare` exits
  `1` on failure; `opendht` always exits `0` and reports failure only in the
  body. Either works, but `success` must always match the real outcome.
- **Reading shell input.** `STUNMESH_KEY` and `STUNMESH_VALUE` are hex
  strings, so `source /dev/stdin` (as in `cloudflare-shell`) is safe today.
  New plugins should parse stdin line by line with an explicit `case`, as
  `opendht-shell` does, so no input line can ever run as a command.
- **TTL backends.** If values in your backend expire, document that users
  must keep `dedup: false` for the plugin instance, as `opendht` does.

A new plugin goes in its own top-level directory with:

- a Makefile with `build`, `clean`, `test`, `install` and `uninstall`
  targets, using the `PLUGIN` variable (not `APP`) for the binary name;
- a binary named `stunmesh-<name>`;
- `CGO_ENABLED ?= 0` for Go plugins, so the binary is static;
- a `smoke_test.sh` or `go test` suite that runs offline;
- a README with setup, configuration, protocol and dependencies.

## License

Apache-2.0. See [LICENSE](LICENSE).
