package ratelimiter

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	// If new instance is acquired without panics, we consider this test a pass
	New(time.Millisecond * 100)
}
func TestRateLimiter_Acquire(t *testing.T) {
	// Procure a new rate limiter with a interval of 100ms
	r := New(time.Millisecond * 10)
	tcs := []acquireTestCase{
		newAcquireTestCase(10, time.Millisecond*90, time.Millisecond*120),
		newAcquireTestCase(11, time.Millisecond*100, time.Millisecond*130),
		newAcquireTestCase(100, time.Millisecond*900, time.Millisecond*1300),
	}

	for _, tc := range tcs {
		start := time.Now()
		for i := 0; i < tc.count; i++ {
			r.Acquire(func() {})
		}
		end := time.Now()

		delta := end.Sub(start)
		if delta < tc.minDelta {
			t.Fatalf("requests were acquired at too quick of a rate! Total time was %v, should have been over %v", delta, tc.minDelta)
		} else if delta > tc.maxDelta {
			t.Fatalf("requests were acquired at too slow of a rate! Total time was %v, should have been under %v", delta, tc.maxDelta)
		}
	}
}

func TestRateLimiter_Close(t *testing.T) {
	r := New(time.Millisecond * 100)
	// The wait functionality is tested within the waiter tests. This test is to ensure
	// there are no errors/panics when the working pieces are put together.
	r.Acquire(func() {})
}

func newAcquireTestCase(count int, minDelta, maxDelta time.Duration) (a acquireTestCase) {
	a.count = count
	a.minDelta = minDelta
	a.maxDelta = maxDelta
	return
}

type acquireTestCase struct {
	count    int
	minDelta time.Duration
	maxDelta time.Duration
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
