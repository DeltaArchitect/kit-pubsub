package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

// Subscriber defines the interface for subscribing to messages
type Subscriber interface {
	Subscribe(ctx context.Context, subID string, handler func(context.Context, *pubsub.Message)) error
}

// Subscribe subscribes to the specified subscription
func (c *Client) Subscribe(ctx context.Context, subID string, handler func(context.Context, *pubsub.Message)) error {
	sub := c.Client.Subscription(subID)
	err := sub.Receive(ctx, handler)
	if err != nil {
		return fmt.Errorf("failed to receive messages: %w", err)
	}

	return nil
}
