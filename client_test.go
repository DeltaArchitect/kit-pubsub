package pubsub

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/pubsub/pstest"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()

	// Start a fake Pub/Sub server.
	srv := pstest.NewServer()
	defer srv.Close()

	// Connect to the fake server.
	conn, err := grpc.Dial(srv.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc.Dial: %v", err)
	}
	defer conn.Close()

	// Create a client.
	client, err := NewClient(ctx, "project-id", option.WithGRPCConn(conn))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	defer client.Close()

	if client.ProjectID != "project-id" {
		t.Errorf("got project ID %q, want %q", client.ProjectID, "project-id")
	}
}

func TestNewClientFromEnv(t *testing.T) {
	ctx := context.Background()
	os.Setenv("GOOGLE_CLOUD_PROJECT", "env-project-id")
	defer os.Unsetenv("GOOGLE_CLOUD_PROJECT")

	// Start a fake Pub/Sub server.
	srv := pstest.NewServer()
	defer srv.Close()

	// Connect to the fake server.
	conn, err := grpc.Dial(srv.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc.Dial: %v", err)
	}
	defer conn.Close()

	// Create a client with empty project ID.
	client, err := NewClient(ctx, "", option.WithGRPCConn(conn))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	defer client.Close()

	if client.ProjectID != "env-project-id" {
		t.Errorf("got project ID %q, want %q", client.ProjectID, "env-project-id")
	}
}
