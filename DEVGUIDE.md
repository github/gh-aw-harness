# Developer Guide

This guide is the maintainer-oriented reference for `gh-aw-harness`.

## Scope

This repository is a harness for GitHub Agentic Workflows. It currently provides a small Go CLI, baseline CI, copied agent skills, a Copilot smoke workflow, and a placeholder specification that will be refined as the harness contract is defined.

The canonical behavior specification for this repository lives in [specs/harness-spec.md](specs/harness-spec.md).

## Development Environment

Preferred setup:

- GitHub Codespaces or the local dev container in [.devcontainer/devcontainer.json](.devcontainer/devcontainer.json)

The dev container keeps common project tooling available, including Go, Docker, GitHub CLI, and Copilot CLI.

### Local Setup

```bash
git clone https://github.com/github/gh-aw-harness.git
cd gh-aw-harness
make deps
make build
```

### Full Maintainer Setup

```bash
make deps-dev
make test
make build
```

## Daily Workflow

Common commands:

```bash
make build          # build the CLI
make test           # run Go tests
make lint           # run go vet
make fmt            # format Go code
make security-scan  # run gosec and govulncheck
make agent-finish   # maintainer validation flow
```

What `make agent-finish` currently runs:

- `deps-dev`
- `fmt`
- `lint`
- `build`
- `test`
- `security-scan`

Use `make help` to see the full list of maintained targets.

## Repository Layout

```text
/
├── .devcontainer/       # Codespaces and Dev Container setup
├── .github/             # Issue templates and workflows
├── cmd/gh-aw-harness/   # CLI entrypoint
├── pkg/harness/         # Minimal harness package
├── skills/              # Small repo-relevant agent skills
├── specs/               # Harness specification scaffold
└── Makefile             # Build and maintainer commands
```

## Agentic Workflow Smoke Test

The Copilot smoke workflow is checked in as both source and generated YAML:

- [.github/workflows/smoke-copilot.md](.github/workflows/smoke-copilot.md)
- [.github/workflows/smoke-copilot.lock.yml](.github/workflows/smoke-copilot.lock.yml)

After editing the markdown source, regenerate the lock file with `gh aw compile` in an environment that has the `gh-aw` GitHub CLI extension installed.

## Maintainer Guidance

When changing code in this repository:

- keep behavior aligned with [specs/harness-spec.md](specs/harness-spec.md) as it evolves
- prefer small, local packages and targeted tests
- update README or spec text when behavior changes
- keep workflow lock files generated from their markdown sources

Useful references:

- [skills/console-rendering/SKILL.md](skills/console-rendering/SKILL.md)
- [skills/error-messages/SKILL.md](skills/error-messages/SKILL.md)

## Troubleshooting

Common reset flow:

```bash
make clean
make deps
make build
make test
```

If tooling is missing:

- run `make deps-dev` to install maintainer dependencies
- reopen the dev container if local environment drift is suspected

## CI Support

The main CI workflow in [.github/workflows/ci.yml](.github/workflows/ci.yml) runs:

- `go vet`
- `go test -race`
- `go build`
- container smoke testing through `make docker-smoke`

## Release Process

The release process for `gh-aw-harness` is not defined yet. Add release workflow documentation here once the image names, versioning model, and promotion process are finalized.
