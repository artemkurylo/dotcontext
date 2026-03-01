# dotcontext

## What is dotcontext?

dotcontext is an open standard for storing all project documentation alongside your code in a `.context/` directory. It defines how complex, enterprise-grade systems should be covered with structured context so that agent-assisted engineering becomes a reliable, everyday practice — not a novelty.

Plans, user stories, RFCs, ADRs, runbooks, onboarding guides — everything lives in the repo, versioned with git, and accessible to both humans and AI agents.

## Problem

Enterprise systems are complex. Dozens of services, years of accumulated decisions, regulatory constraints, intricate domain models. The knowledge needed to work on these systems is scattered across wikis, Notion pages, Google Docs, Confluence spaces, Slack threads, and tribal memory. This creates critical problems:

- **Context loss** — when engineers switch projects or leave, critical knowledge disappears with them
- **Stale docs** — documentation hosted outside the repo drifts out of sync with the code it describes
- **Agent-unfriendly** — AI coding agents can read your codebase but not your Confluence, so they operate with partial understanding of complex systems
- **Review gaps** — code gets reviewed in PRs, but the architectural decisions and business reasoning behind it don't
- **Onboarding friction** — new engineers on large systems spend weeks piecing together context that should be at their fingertips

## Solution

Put everything in `.context/` and enforce it. Documentation becomes a first-class citizen of the codebase — reviewed in PRs, versioned in git, and always in sync with the code it describes.

## Core Principles

1. **Code and context together** — if it explains the code, it belongs next to the code
2. **Git-native** — all documentation is versioned, diffable, and reviewable in PRs
3. **Convention over configuration** — a standard directory structure so every project feels familiar
4. **Agent-accessible** — AI agents can read and reason over the full project context
5. **Enforced, not optional** — linting and CI checks ensure documentation stays current

## Directory Structure

```
.context/
  CONTEXT.md        # This file. Project overview and orientation.
  plans/            # Implementation plans and roadmaps
  stories/          # User stories and requirements
  rfcs/             # Requests for comments on proposed changes
  adrs/             # Architecture decision records
  runbooks/         # Operational procedures
  guides/           # Onboarding and how-to guides
```

## Who is this for?

dotcontext targets teams building complex, enterprise-grade systems — the kind where context is the bottleneck, not code.

- **Enterprise engineering teams** managing large codebases with deep domain complexity, regulatory requirements, and long-lived services
- **Platform and infrastructure teams** where architectural decisions have lasting consequences and must be understood years later
- **Organizations adopting AI-assisted engineering** that need agents to operate with full system understanding, not just file-level code context
- **Any team** where the cost of lost context — wrong decisions, repeated mistakes, slow onboarding — is too high to tolerate

## License

dotcontext is fully open source. Everyone is welcome to use, contribute to, and build upon it.

- **Specification and documentation** — licensed under [CC-BY-SA 4.0](https://creativecommons.org/licenses/by-sa/4.0/). You can freely adopt, share, and adapt the standard. You must give credit to dotcontext, and any derivative specifications must be shared under the same terms.
- **Tooling and code** — licensed under [MPL 2.0](https://www.mozilla.org/en-US/MPL/2.0/). You can use dotcontext tooling in any project, including proprietary ones. Modifications to dotcontext source files themselves must be shared back under the same license.
- **The dotcontext name** is a trademark. You can use the standard freely, but you cannot rebrand or misrepresent it as your own.

## Status

This project is in the early design phase. We are building the convention, the CLI tooling, and the CI integrations to make `.context/` the standard way to keep project knowledge alive in complex systems.
