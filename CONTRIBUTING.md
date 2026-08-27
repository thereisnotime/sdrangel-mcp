# Contributing

## Setup

```sh
git clone git@github.com:thereisnotime/sdrangel-mcp.git
cd sdrangel-mcp
cp .env.example .env
just build
```

## Before submitting a PR

```sh
just check   # tidy + fmt + vet + build + test
just lint    # golangci-lint
```

## Commit style

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: add VOR channel tool
fix: handle empty response from spectrum server
chore: update go-sdk to v1.7.0
```

## Adding a new tool group

1. Add client methods in `internal/sdrangel/<group>.go`.
2. Add types to `internal/sdrangel/types.go` if needed.
3. Create `internal/tools/<group>.go` with a `registerXxxTools` function.
4. Call it from `internal/tools/server.go`.
5. Update the tool table in `README.md` and the matching section in `docs/tools-reference.md`.
