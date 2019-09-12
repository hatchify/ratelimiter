package ratelimiter

import "testing"

func TestNewRequest(t *testing.T) {
	// If new request is acquired without panics, we consider this test a pass
	newRequest(func() {})

}
