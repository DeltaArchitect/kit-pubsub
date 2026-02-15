# kit-pubsub

[![Go Report Card](https://goreportcard.com/badge/github.com/DeltaArchitect/kit-pubsub)](https://goreportcard.com/report/github.com/DeltaArchitect/kit-pubsub)
[![codecov](https://codecov.io/gh/DeltaArchitect/kit-pubsub/branch/main/graph/badge.svg)](https://codecov.io/gh/DeltaArchitect/kit-pubsub)
[![Go Reference](https://pkg.go.dev/badge/github.com/DeltaArchitect/kit-pubsub.svg)](https://pkg.go.dev/github.com/DeltaArchitect/kit-pubsub)

A simplified, testable Pub/Sub V2 client wrapper for Go.

## Installation

```bash
go get github.com/DeltaArchitect/kit-pubsub
```

## Usage

### Initialization

```go
import "github.com/DeltaArchitect/kit-pubsub"

ctx := context.Background()
client, err := pubsub.NewClient(ctx, "your-project-id")
if err != nil {
    // handle error
}
defer client.Close()
```

### Publishing Messages

```go
topicID := "your-topic"
data := []byte("Hello, Pub/Sub!")
id, err := client.Publish(ctx, topicID, data)
if err != nil {
    // handle error
}
fmt.Printf("Published message with ID: %s\n", id)
```

### Subscribing to Messages

```go
subID := "your-subscription"
err := client.Subscribe(ctx, subID, func(ctx context.Context, msg *pubsub.Message) {
    fmt.Printf("Received message: %s\n", string(msg.Data))
    msg.Ack()
})
if err != nil {
    // handle error
}
```

## Testing

This library includes helpers for testing your application code that uses Pub/Sub.

You can use `pubsub.NewClient` with `option.WithGRPCConn(conn)` to connect to a fake Pub/Sub server (like `pstest`).


