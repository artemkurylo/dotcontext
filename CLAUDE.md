# CLAUDE.md — AI Workflow Guide

This file tells AI agents how to work with the dotcontext project. Read it before making changes.

## Start here

1. Read `.context/CONTEXT.md` for project overview, architecture, and tech stack
2. Check `.context/adrs/` for past architectural decisions before proposing new ones
3. Follow the workflow below for all tasks

## Project structure

```
cmd/dotcontext/          # CLI entry point (Cobra root command)
internal/
  init/                  # `dotcontext init` — scaffolds .context/ directory
  check/                 # `dotcontext check` — validates .context/ structure
  templates/             # Embedded templates (Go embed, *.tmpl files)
.context/                # Project's own context directory
  CONTEXT.md             # Project overview and orientation
  plans/                 # Implementation plans
  stories/               # User stories
  rfcs/                  # Requests for comments
  adrs/                  # Architecture decision records
  runbooks/              # Operational procedures
  guides/                # Onboarding and how-to guides
```

## Build and test

```bash
make build               # Build binary to bin/dotcontext
make test                # Run all tests
make lint                # Run golangci-lint
make clean               # Remove build artifacts
```

## Task-type lookup

| Task | Read first | Write to |
|------|-----------|----------|
| Plan a feature | `.context/plans/` | `.context/plans/` |
| Record a decision | `.context/adrs/` | `.context/adrs/` |
| Propose a change | `.context/rfcs/` | `.context/rfcs/` |
| Fix an incident | `.context/runbooks/` | `.context/runbooks/` |
| Onboard someone | `.context/guides/` | `.context/guides/` |

## Workflow

1. **Read context** — start with `.context/CONTEXT.md`, then check the relevant subdirectory for your task type
2. **Check decisions** — review `.context/adrs/` so you don't contradict existing architectural choices
3. **Follow conventions** — match the code style, naming, and patterns already in the codebase
4. **Update docs** — if your change affects architecture, decisions, or operations, update the relevant `.context/` files
5. **Validate** — run `make test` and `make lint` before finishing

## Code conventions

- **License header** — every `.go` file starts with the MPL 2.0 header comment
- **Package naming** — use short, lowercase names; `internal/` packages are not exported
- **Error wrapping** — use `fmt.Errorf("context: %w", err)` to wrap errors with context
- **Templates** — placed in `internal/templates/` as `*.tmpl` files; the `//go:embed *.tmpl` glob picks them up automatically
- **Tests** — use Go's standard `testing` package; table-driven tests where appropriate; test files live next to the code they test
