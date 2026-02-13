package pubsub

import "errors"

var (
	ErrEmptyProjectID = errors.New("project ID is required")
	ErrTopicNotFound  = errors.New("topic not found")
	ErrSubNotFound    = errors.New("subscription not found")
)
