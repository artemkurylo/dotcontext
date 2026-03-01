# Adopting dotcontext

A step-by-step workflow for adding dotcontext to an existing project.

You have a project. The knowledge about it is scattered — some in Confluence, some in Google Docs, some in Slack threads, some in people's heads. You want to fix that. Here's how.

---

## Phase 1: Initialize

Install the dotcontext CLI and scaffold the directory structure.

```bash
npm install -g dotcontext
```

Then, in the root of your repository:

```bash
dotcontext init
```

This creates the full `.context/` structure:

```
.context/
  CONTEXT.md          # Project overview (template, ready to fill in)
  plans/
  stories/
  rfcs/
  adrs/
  runbooks/
  guides/
```

The CLI generates a starter `CONTEXT.md` with sections to fill out — what the project is, architecture, tech stack, team, key decisions. Commit this scaffolding as your first step.

## Phase 2: Generate context with AI

Open your project in Claude Code and run:

```
dotcontext generate
```

Claude scans your codebase — file structure, dependencies, configuration, code patterns — and generates an initial set of context artifacts:

- **CONTEXT.md** — filled in with detected architecture, tech stack, service boundaries, and project structure
- **ADRs** — inferred from code patterns (e.g., "this project uses PostgreSQL with Prisma", "authentication is handled via JWT")
- **Guides** — a developer setup guide generated from your package.json, Dockerfile, Makefile, or similar entry points

Review what was generated, edit as needed, and commit. This gets you from zero to a meaningful `.context/` in minutes instead of days.

## Phase 3: Import from external sources

Most teams already have documentation — it's just in the wrong place. Use MCP integrations to pull it into `.context/`:

```
dotcontext import --from confluence
dotcontext import --from notion
dotcontext import --from jira
```

The CLI connects to your existing tools via MCP (Model Context Protocol), reads your documentation, and converts it to the dotcontext format:

- **Confluence pages** tagged as ADRs or design docs → `.context/adrs/`
- **Notion runbooks** → `.context/runbooks/`
- **Jira epics and stories** → `.context/stories/`
- **Google Docs RFCs** → `.context/rfcs/`

Each imported document is converted to markdown, placed in the correct directory, and formatted to match the dotcontext templates. You review and commit — nothing is modified at the source.

## Phase 4: Fill in the gaps

AI generation and imports won't capture everything. Some context only exists in people's heads. After phases 2 and 3, review what's missing:

**CONTEXT.md** — Does it accurately describe the system? Add business context, team structure, and key decisions that code analysis can't detect.

**ADRs** — What critical decisions are undocumented? Start with the questions people ask most often: "Why do we use X instead of Y?" Each answer is an ADR.

```markdown
# ADR-001: [Title of decision]

## Status
Accepted | Superseded | Deprecated

## Context
What situation or problem prompted this decision?

## Decision
What did we decide?

## Consequences
What are the trade-offs? What becomes easier or harder?
```

**RFCs** — Any active proposals or design discussions happening in Slack or meetings? Move them here.

```markdown
# RFC-001: [Title]

## Status
Draft | In Review | Accepted | Rejected | Implemented

## Summary
One paragraph: what are you proposing?

## Motivation
Why is this change needed? What problem does it solve?

## Proposal
Detailed description of the proposed change.

## Alternatives considered
What else did you consider, and why was it rejected?

## Open questions
What still needs to be resolved?
```

**Runbooks** — Deployment procedures, incident response, database migrations, on-call handoff. These are the documents people reach for under pressure — they must be accurate and findable.

```markdown
# [Runbook title]

## When to use this
What situation triggers this runbook?

## Prerequisites
Access, permissions, tools needed.

## Steps
1. Step one
2. Step two
3. ...

## Rollback
How to undo if something goes wrong.

## Contacts
Who to escalate to.
```

**Guides** — Developer setup, onboarding, domain glossary. If a new engineer would struggle without it, write it down.

## Phase 5: Start working in dotcontext

With the foundation in place, shift to using `.context/` for active work:

**Plans** (`.context/plans/`) — implementation plans for upcoming work. What are we building, how are we breaking it down, what's the approach? Write plans here, review them in PRs, and let AI agents reference them during implementation.

**Stories** (`.context/stories/`) — user stories and requirements for work in progress. Acceptance criteria, edge cases, constraints. When an agent picks up a story, it reads the full context — not just a Jira title.

This is where dotcontext becomes a daily workflow, not a one-time migration.

## Phase 6: Enforce

Adoption fails when it's optional. Make dotcontext part of how your team works:

```
dotcontext check
```

The CLI validates your `.context/` directory:
- `CONTEXT.md` exists and is not empty
- Required sections are present
- ADR and RFC numbering is consistent
- No broken internal links

Add it to CI so it runs on every PR. Combine with:

- **PR template** — a checkbox: "Does this change need a CONTEXT.md update, a new ADR, or a runbook change?"
- **Team agreement** — significant technical decisions get an ADR. New features get a plan. Non-negotiable.
- **Review culture** — review `.context/` changes with the same rigor as code

---

## Summary

| Phase | What happens | How |
|-------|-------------|-----|
| 1. Initialize | Scaffold `.context/` structure | `dotcontext init` |
| 2. Generate | AI scans codebase, generates initial artifacts | `dotcontext generate` (Claude Code) |
| 3. Import | Pull existing docs from external tools | `dotcontext import --from <source>` (MCP) |
| 4. Fill gaps | Manually add context that only humans know | Edit markdown files |
| 5. Active use | Plans and stories for daily work | Write in `.context/`, review in PRs |
| 6. Enforce | CI checks and team agreements | `dotcontext check` |

---

## Common questions

**Do I need to migrate everything at once?**
No. Phase 1-2 take minutes. Phase 3 takes an hour. Phase 4 is ongoing — start with the highest-value gaps and fill in over time.

**What about docs that change frequently?**
That's the point. They live in git, so changes are tracked, diffed, and reviewed — just like code.

**Should I delete the old wiki/Confluence pages?**
After migration, mark them as deprecated with a link to the `.context/` version. Remove them once the team has fully transitioned.

**What if my project is a monorepo with multiple services?**
Each service can have its own `.context/` directory. The root `.context/` covers the system as a whole.

**Do I need Claude Code to use dotcontext?**
No. The standard works with any editor and any AI agent. Claude Code integration makes Phase 2 (generation) faster, but the directory structure and conventions are tool-agnostic.
