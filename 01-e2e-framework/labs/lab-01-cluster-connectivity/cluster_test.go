package clusterconnectivity

import (
	"context"
	"os"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/features"
)

var testenv env.Environment

func TestMain(m *testing.M) {
	testenv = env.New()
	os.Exit(testenv.Run(m))
}

func TestClusterConnectivity(t *testing.T) {
	feature := features.New("cluster connectivity").
		WithLabel("type", "connectivity").
		Assess("list namespaces and find kube-system", func(ctx context.Context, t *testing.T, cfg *envconf.Config) context.Context {
			var namespaces corev1.NamespaceList

			if err := cfg.Client().Resources().List(ctx, &namespaces); err != nil {
				t.Fatalf("failed to list namespaces: %v", err)
			}

			t.Logf("connected successfully; discovered %d namespaces", len(namespaces.Items))

			foundKubeSystem := false
			for _, ns := range namespaces.Items {
				t.Logf("namespace: %s", ns.Name)
				if ns.Name == "kube-system" {
					foundKubeSystem = true
				}
			}

			if !foundKubeSystem {
				t.Fatal("kube-system namespace was not found")
			}

			return ctx
		}).
		Feature()

	testenv.Test(t, feature)
}
