package pubsub

import (
	"context"
	"testing"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestPublish(t *testing.T) {
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
	client, err := pubsub.NewClient(ctx, "project-id", option.WithGRPCConn(conn))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	defer client.Close()

	c := &Client{
		Client:    client,
		ProjectID: "project-id",
	}

	topicID := "test-topic"
	_, err = c.Client.CreateTopic(ctx, topicID)
	if err != nil {
		t.Fatalf("CreateTopic: %v", err)
	}

	data := []byte("hello world")
	id, err := c.Publish(ctx, topicID, data)
	if err != nil {
		t.Fatalf("Publish: %v", err)
	}

	if id == "" {
		t.Errorf("got empty message ID")
	}
}

func TestPublishError(t *testing.T) {
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
	client, err := pubsub.NewClient(ctx, "project-id", option.WithGRPCConn(conn))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	defer client.Close()

	c := &Client{
		Client:    client,
		ProjectID: "project-id",
	}

	// Create a context that is already canceled to force Publish to fail
	ctxCanceled, cancel := context.WithCancel(ctx)
	cancel()

	_, err = c.Publish(ctxCanceled, "test-topic", []byte("data"))
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestPublishTopicNotFound(t *testing.T) {
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
	client, err := pubsub.NewClient(ctx, "project-id", option.WithGRPCConn(conn))
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	defer client.Close()

	c := &Client{
		Client:    client,
		ProjectID: "project-id",
	}

	// Publish to a non-existent topic
	_, err = c.Publish(ctx, "non-existent-topic", []byte("data"))
	if err == nil {
		t.Error("expected error, got nil")
	}
}
