# Kubernetes Testing Framework Learning

Hands-on learning repository for Kubernetes testing approaches.

## Learning roadmap

1. **Kubernetes SIG E2E Framework** — current focus
2. **Kubernetes Test Framework (KTF)** — later
3. **Operator-specific test harnesses** — later

The repository is intentionally built one lab at a time so each concept can be understood, executed, broken, debugged, and improved before moving forward.

## Current phase: E2E Framework

Initial progression:

- Go testing fundamentals
- E2E Framework architecture
- Connect to an existing Kubernetes cluster
- Environment lifecycle: setup and finish
- Features and assessments
- Kubernetes client usage
- Resource creation and cleanup
- Kind-based ephemeral clusters
- Workload, networking, storage, scheduling, and security tests
- CI automation

## Repository layout

```text
.
├── 01-e2e-framework/
│   └── labs/
├── 02-ktf/
├── 03-operator-testing/
├── docs/
├── AGENTS.md
└── README.md
```

Only the E2E section is active now. The other sections are roadmap placeholders.

## Upstream project

https://github.com/kubernetes-sigs/e2e-framework

## Learning principle

Each lab should explain what problem the test solves, the framework concepts introduced, how the code works, how to run it, successful behavior, a deliberate failure exercise, debugging, cleanup, and lessons learned.
