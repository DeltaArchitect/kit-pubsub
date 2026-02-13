package pubsub

import (
	"context"
	"sync"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestSubscribe(t *testing.T) {
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
	topic, err := c.Client.CreateTopic(ctx, topicID)
	if err != nil {
		t.Fatalf("CreateTopic: %v", err)
	}

	subID := "test-sub"
	_, err = c.Client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{
		Topic: topic,
	})
	if err != nil {
		t.Fatalf("CreateSubscription: %v", err)
	}

	// Publish a message
	_, err = c.Publish(ctx, topicID, []byte("hello"))
	if err != nil {
		t.Fatalf("Publish: %v", err)
	}

	// Subscribe
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		err := c.Subscribe(ctx, subID, func(ctx context.Context, msg *pubsub.Message) {
			msg.Ack()
			if string(msg.Data) != "hello" {
				t.Errorf("got message %q, want %q", string(msg.Data), "hello")
			}
			wg.Done()
		})
		if err != nil {
			// In a real scenario we might handle this differently, but for the test,
			// if Subscribe returns before processing, checking err is good.
			// However, Subscribe blocks. So valid use is checking context cancellation or error return.
			// Here we just log.
			t.Logf("Subscribe returned: %v", err)
		}
	}()

	// Wait for message to be processed or timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// success
	case <-time.After(5 * time.Second):
		t.Fatal("timeout waiting for message")
	}
}

func TestSubscribeError(t *testing.T) {
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

	// Subscribe to a non-existent subscription
	err = c.Subscribe(ctx, "non-existent-sub", func(ctx context.Context, msg *pubsub.Message) {})
	if err == nil {
		t.Error("expected error, got nil")
	}
}
