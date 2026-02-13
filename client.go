package pubsub

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

// Client wraps the Google Cloud Pub/Sub client
type Client struct {
	*pubsub.Client
	ProjectID string
}

// NewClient creates a new Pub/Sub client
func NewClient(ctx context.Context, projectID string, opts ...option.ClientOption) (*Client, error) {
	if projectID == "" {
		projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	}

	if projectID == "" {
		return nil, fmt.Errorf("project ID is required")
	}

	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	return &Client{
		Client:    client,
		ProjectID: projectID,
	}, nil
}

// Close closes the Pub/Sub client
func (c *Client) Close() error {
	return c.Client.Close()
}
