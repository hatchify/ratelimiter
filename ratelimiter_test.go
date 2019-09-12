package ratelimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	// If new rate limiter is acquired without panics, we consider this test a pass
	New(time.Millisecond * 100)
}
func TestRateLimiter_Acquire(t *testing.T) {
	// Procure a new rate limiter with a interval of 100ms
	r := New(time.Millisecond * 10)
	tcs := []acquireTestCase{
		newAcquireTestCase(10, time.Millisecond*90),
		newAcquireTestCase(11, time.Millisecond*100),
		newAcquireTestCase(100, time.Millisecond*900),
	}

	for _, tc := range tcs {
		start := time.Now()
		for i := 0; i < tc.count; i++ {
			r.Acquire(func() {})
		}
		end := time.Now()

		duration := end.Sub(start)
		if duration < tc.minDuration {
			t.Fatalf("requests were acquired at too quick of a rate! Total time was %v, should have been over %v", duration, tc.minDuration)
		}
	}
}

func TestRateLimiter_Close(t *testing.T) {
	r := New(time.Millisecond * 100)
	// The wait functionality is tested within the waiter tests. This test is to ensure
	// there are no errors/panics when the working pieces are put together.
	r.Acquire(func() {})
}

func newAcquireTestCase(count int, minDuration time.Duration) (a acquireTestCase) {
	a.count = count
	a.minDuration = minDuration
	return
}

type acquireTestCase struct {
	count       int
	minDuration time.Duration
}

func ExampleNew() {
	// Initialize rate limiter
	r := New(time.Second)
	fmt.Printf("RateLimiter (%v) is ready to use!\n", r)
}

func ExampleRateLimiter_Acquire() {
	// Initialize rate limiter
	r := New(time.Second)
	// Acquire a new request
	r.Acquire(func() {
		// Do stuff here
	})
}

func ExampleRateLimiter_Close() {
	// Initialize rate limiter
	r := New(time.Second)

	// Do lots of great tasks here

	// Close rate limiter
	r.Close()
}
