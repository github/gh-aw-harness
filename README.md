# gh-aw-harness

Harness repository for [GitHub Agentic Workflows](https://github.com/github/gh-aw). This repo is a small Go-based starting point for exercising Agentic Workflow behavior, smoke tests, and repository automation before fuller harness behavior is implemented.

## Quick Start

Build the CLI:

```bash
make build
```

Run it locally:

```bash
./bin/gh-aw-harness --version
```

Run tests:

```bash
make test
```

## Current Scope

This repository is intentionally minimal. It currently includes:

- a stub Go command in [cmd/gh-aw-harness](cmd/gh-aw-harness)
- a tiny package in [pkg/harness](pkg/harness)
- baseline CI in [.github/workflows/ci.yml](.github/workflows/ci.yml)
- a Copilot Agentic Workflow smoke test in [.github/workflows/smoke-copilot.md](.github/workflows/smoke-copilot.md)
- copied agent skills in [skills](skills)
- a placeholder harness specification in [specs/harness-spec.md](specs/harness-spec.md)

The harness contract and production behavior still need to be defined in the spec and implemented in the CLI.

## Usage

```bash
gh-aw-harness [flags]
```

Flags:

- `--version` - print the current build version and exit

## Container

Build the container image:

```bash
make docker-build
```

Run the container smoke test:

```bash
make docker-smoke
```

## Development

Common commands:

```bash
make build          # build the CLI
make test           # run Go tests
make lint           # run go vet
make fmt            # format Go code
make security-scan  # run gosec and govulncheck
```

Use `make help` for the full list of maintained targets.

## Agentic Workflow Smoke Test

The Copilot smoke workflow is kept in two forms:

- [.github/workflows/smoke-copilot.md](.github/workflows/smoke-copilot.md) - Agentic Workflow source
- [.github/workflows/smoke-copilot.lock.yml](.github/workflows/smoke-copilot.lock.yml) - compiled workflow

After editing the markdown workflow, regenerate the lock file with `gh aw compile` in an environment that has the `gh-aw` GitHub CLI extension installed.

## Specification

See [specs/harness-spec.md](specs/harness-spec.md). It is currently a scaffold and must be rewritten as the harness contract is defined.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md) for development setup and contribution guidelines.

## Support

See [SUPPORT.md](SUPPORT.md) for help, issue reporting, and support scope.

## Code of Conduct

See [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md).

## Security

See [SECURITY.md](SECURITY.md) for vulnerability reporting instructions.

## License

See [LICENSE](LICENSE) for details.
