package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// Publisher defines the interface for publishing messages
type Publisher interface {
	Publish(ctx context.Context, topicID string, data []byte) (string, error)
}

// Publish publishes a message to the specified topic
func (c *Client) Publish(ctx context.Context, topicID string, data []byte) (string, error) {
	t := c.Client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	id, err := result.Get(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to publish message: %w", err)
	}

	return id, nil
}
