# ADR-001: Use Go with Cobra for the CLI

## Status

Accepted

## Context

We need to build a CLI tool (`dotcontext`) that scaffolds directory structures, validates files, and eventually integrates with AI services and MCP connectors. The CLI will be distributed as an open source tool that anyone can install.

Key requirements:
- Easy to install — minimal friction for first-time users
- No runtime dependencies — users shouldn't need Node.js, Python, or any other runtime pre-installed
- Fast startup — CLI tools that take >200ms to start feel sluggish
- Cross-platform — must work on macOS, Linux, and Windows
- Contributor-friendly — the language should be accessible to a wide range of open source contributors

## Decision

Use **Go** as the language and **Cobra** as the CLI framework.

**Go** compiles to a single static binary with no runtime dependencies. It cross-compiles trivially for all major platforms. The language is simple and readable, making it accessible to contributors who may not be Go experts.

**Cobra** is the de facto standard CLI framework in Go. It powers kubectl, gh (GitHub CLI), docker, and hugo. It provides subcommand routing, flag parsing, help generation, and shell completions out of the box.

## Alternatives Considered

**TypeScript / Node.js**
- Pro: largest ecosystem, many team members already know it
- Con: requires Node.js runtime installed. `npm install -g` can conflict with system Node versions, permission issues, nvm complexity. Startup time is noticeably slower than compiled Go.
- Con: distributing a CLI that requires a runtime is friction for adoption

**Rust**
- Pro: single binary like Go, even faster execution, strong type system
- Con: steeper learning curve reduces the contributor pool
- Con: compile times are significantly longer, slowing development iteration
- Con: the ecosystem for CLI tooling is less mature than Go's (clap is good but Cobra's ecosystem is richer)

## Consequences

- **Single binary distribution** — users download one file and it works. No package managers, no runtimes.
- **GitHub Releases with cross-compilation** — GoReleaser can produce binaries for all platforms in CI
- **Fast startup** — Go binaries start in single-digit milliseconds
- **Go ecosystem** — we get access to the Go standard library for file I/O, template rendering, and testing without external dependencies
- **Contributor accessibility** — Go's simplicity means contributors can be productive quickly, even if Go isn't their primary language
- **Future AI integration** — Go has solid HTTP client support for calling Claude API, and can shell out to or embed MCP tooling
