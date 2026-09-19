# Lab 01 — Cluster Connectivity

## Objective

Learn the smallest useful Kubernetes SIG E2E Framework test against an **existing Kubernetes cluster**.

This lab intentionally does not create or modify Kubernetes resources. It only connects to the cluster, lists namespaces, and verifies that the standard `kube-system` namespace exists.

## Concepts introduced

- Go's `testing` package
- `TestMain`
- `env.Environment`
- `env.New()`
- `Environment.Run()`
- `features.New()`
- `Assess()`
- `envconf.Config`
- the framework Kubernetes client
- read-only Kubernetes API interaction

## Architecture

```text
go test
   |
   v
TestMain
   |
   v
env.New()
   |
   v
Feature -> Assessment
   |
   v
envconf.Config / klient
   |
   v
kubeconfig
   |
   v
Kubernetes API Server
```

## Prerequisites

You need:

- Go installed
- `kubectl` installed
- access to a Kubernetes cluster
- a working kubeconfig

Check them first:

```bash
go version
kubectl version --client
kubectl config current-context
kubectl cluster-info
kubectl get namespaces
```

Do not continue until `kubectl get namespaces` succeeds.

## Project setup

From this lab directory:

```bash
go mod tidy
```

This downloads the E2E Framework and Kubernetes API dependencies declared by the lab.

## Run the test

```bash
go test -v .
```

The test should connect using your normal Kubernetes client configuration and print the namespaces it discovers.

A successful run ends with output similar to:

```text
--- PASS: TestClusterConnectivity
PASS
```

The exact framework log output may vary by framework version.

## What the test does

The suite creates an E2E `Environment` in `TestMain`.

The test then defines one Feature named `cluster connectivity`.

That Feature contains one Assessment:

```text
list namespaces and find kube-system
```

The assessment obtains the Kubernetes client from the framework configuration and asks the API server for a `NamespaceList`. It fails if the API request fails or if `kube-system` is not present.

## Failure exercise 1 — invalid kubeconfig

Keep your working kubeconfig safe. In a shell used only for this exercise, point `KUBECONFIG` at a nonexistent file and run the test.

```bash
export KUBECONFIG=/tmp/nonexistent-kubeconfig
go test -v .
```

Observe where configuration/client initialization fails.

Restore your configuration afterward, for example:

```bash
unset KUBECONFIG
kubectl get namespaces
```

If you normally set `KUBECONFIG` explicitly, restore its original value instead of unsetting it.

## Failure exercise 2 — change the assertion

Temporarily change:

```go
if ns.Name == "kube-system" {
```

to a namespace that does not exist, such as:

```go
if ns.Name == "namespace-that-should-not-exist" {
```

Run:

```bash
go test -v .
```

The Kubernetes connection still succeeds, but the **assessment** fails. This distinction is important: connectivity failure and assertion failure are different classes of test failure.

Revert the change after the exercise.

## Cleanup

This lab is read-only and creates no Kubernetes resources, so there is no cluster cleanup.

## Lessons to verify before Lab 02

You should be able to explain:

1. Why `TestMain` exists.
2. What `env.New()` creates.
3. What a Feature represents.
4. What an Assessment represents.
5. Where `envconf.Config` comes from.
6. How the framework obtains a Kubernetes client.
7. The difference between an API/connectivity failure and a failed assertion.

Do not move to Lab 02 until these concepts are comfortable.
