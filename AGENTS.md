# AGENTS.md

## Purpose

This repository is a hands-on learning environment for Kubernetes testing frameworks.

The current focus is the Kubernetes SIG E2E Framework. Kubernetes Test Framework (KTF) and operator-specific test harnesses will be studied later.

## Learning rules

- Keep examples generic and independent of company or production applications.
- Teach concepts progressively rather than generating a large finished framework at once.
- Prefer small, runnable examples.
- Explain important Go and Kubernetes concepts used by each test.
- Include commands, expected behavior, deliberate failure exercises where useful, debugging guidance, and cleanup.
- Avoid abstractions until the underlying framework behavior has been demonstrated.
- Prefer upstream Kubernetes and framework APIs and documentation.

## Repository roadmap

### 01-e2e-framework

Current track: environment configuration, TestMain, features, assessments, Kubernetes clients, resource lifecycle, Kind integration, labels/filtering, networking, storage, scheduling, security, failure testing, reusable helpers, and CI.

### 02-ktf

Reserved for Kubernetes Test Framework learning after the E2E Framework foundations.

### 03-operator-testing

Reserved for operator-specific test harnesses after the earlier foundations.

## Lab conventions

Use sequential names such as `lab-01-cluster-connectivity`, `lab-02-resource-lifecycle`, and `lab-03-workload-testing`.

Each lab should normally contain source code and a README covering prerequisites, concepts, execution, expected results, failure exercises, cleanup, and lessons learned.

## Scope

This is an educational repository. Prefer generic resources such as namespaces, simple nginx workloads, Services, ConfigMaps, PVCs, and standard Kubernetes APIs unless a lab specifically studies another component.
