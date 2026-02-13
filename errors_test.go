package pubsub

import (
	"errors"
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want string
	}{
		{
			name: "ErrEmptyProjectID",
			err:  ErrEmptyProjectID,
			want: "project ID is required",
		},
		{
			name: "ErrTopicNotFound",
			err:  ErrTopicNotFound,
			want: "topic not found",
		},
		{
			name: "ErrSubNotFound",
			err:  ErrSubNotFound,
			want: "subscription not found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestNewClientError(t *testing.T) {
	_, err := NewClient(nil, "")
	if !errors.Is(err, ErrEmptyProjectID) && err.Error() != "project ID is required" {
		t.Errorf("got %v, want %v", err, ErrEmptyProjectID)
	}
}
