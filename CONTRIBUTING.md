# Contributing to GitHub Agentic Workflows Harness

Thank you for your interest in contributing to `gh-aw-harness`. This repository is a harness for GitHub Agentic Workflows and is developed by a core team using agentic development.

Important: if you are not part of the core team, please do not create pull requests directly. Instead, create a detailed agentic plan in an issue, discuss it with the team, and a core team member will create and implement the PR using agents.

This document describes the contribution process for non-core team members.

## Prerequisites

Generic development environments with manually installed tools are not the supported path. This project is designed to be developed inside a Dev Container or GitHub Codespace, which automatically configures the required tools and runtimes.

The recommended way to set up a development environment is to use the provided [Dev Container](.devcontainer/devcontainer.json):

- GitHub Codespaces: open this repository in a Codespace. The environment is preconfigured with Go, Docker, and the GitHub CLI.
- VS Code Dev Container: install the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers), then open the repository folder and choose **Reopen in Container**.

The Dev Container installs required dependencies and runs `make deps` automatically on creation. If you encounter errors about Go or other tool versions when running `make deps` or build commands, switch to the Dev Container or Codespace environment.

## How Development Works

GitHub Agentic Workflows projects are developed by a core team using agentic development, primarily GitHub Copilot coding agent and local coding agents. This means:

- Core team members use agents to create and manage pull requests.
- CI runs quality checks on pull requests.
- Community contributors propose changes through detailed agentic plans in issues.
- Traditional pull requests from non-core members are not enabled.

## Quick Start for Community Contributors

If you are not part of the core team, do not create pull requests directly. Instead, craft a detailed agentic plan in an issue and a core team member will pick it up and implement it using agents.

### Step 1: Analyze with an Agent

Before filing a contribution request, use an agent to:

- scan the source code to identify root causes for bugs
- analyze execution patterns and trace the issue
- research similar issues and solutions
- propose specific fixes with code examples
- create a comprehensive plan for the changes needed

### Step 2: Open an Issue with Your Agentic Plan

Create an issue with your detailed agentic plan:

- describe what you want to contribute
- include your agent's analysis and findings for bugs
- explain the use case and expected behavior
- provide a complete, step-by-step plan for the agent to follow
- include specific implementation details and examples
- tag with appropriate labels

### Step 3: Discuss and Refine with the Team

Once you have opened the issue:

1. A core team member reviews your plan and may ask clarifying questions.
2. You iterate on the implementation approach based on team feedback.
3. A core team member signals when they will pick it up.

### Step 4: A Core Team Member Implements the PR

A core team member will:

- take your agentic plan and use a coding agent to implement it
- read relevant documentation and specifications
- make code changes following established patterns
- run the project validation checks
- create a PR and handle review feedback

You do not create or own the PR. The core team member does, using agents as their implementation tool.

## Project Structure

This structure is useful context when writing your agentic plan:

```text
/
├── .devcontainer/       # Codespaces and Dev Container setup
├── .github/             # Issue templates and workflows
├── cmd/gh-aw-harness/   # Main CLI application
├── pkg/harness/         # Minimal harness package
├── skills/              # Specialized knowledge for agents
├── specs/               # Harness specification scaffold
├── DEVGUIDE.md          # Maintainer and agent reference guide
└── Makefile             # Build automation and validation targets
```

## Dependency License Policy

This project uses an MIT license and only accepts dependencies with compatible licenses.

Allowed licenses include:

- MIT
- Apache-2.0
- BSD-2-Clause and BSD-3-Clause
- ISC

Disallowed licenses include:

- GPL, LGPL, and AGPL
- SSPL
- proprietary or commercial licenses requiring payment or special terms

## Spam Prevention

Be nice, do not spam. The project maintainers reserve the right to clean up spam, unsolicited promotions, or off-topic content as needed to keep discussions focused and valuable.

This includes:

- repeated identical or similar comments across issues or pull requests
- unsolicited promotional content or advertisements
- off-topic comments that do not contribute to the discussion
- automated bot comments without prior approval

## Community

- Join the `#continuous-ai` channel in the [GitHub Next Discord](https://gh.io/next-discord)
- Participate in discussions on GitHub issues
- Collaborate by crafting high-quality agentic plans for the core team to implement

## Code of Conduct

This project follows the GitHub Community Guidelines. Please be respectful and inclusive in all interactions.

## Getting Help

- For bugs or features: open a GitHub issue with a detailed agentic plan.
- For questions: ask in issues, discussions, or Discord.
- For examples: look at existing issues and PRs created by core team members.

The [Development Guide](DEVGUIDE.md) is the reference guide used by core team members and their agents.

Thank you for contributing to GitHub Agentic Workflows Harness.
