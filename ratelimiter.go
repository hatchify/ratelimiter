package ratelimiter

import (
	"sync"
	"time"

	"github.com/Hatch1fy/errors"
)

// New will return a new RateLimiter
func New(interval time.Duration) *RateLimiter {
	var r RateLimiter
	// Set the rate interval
	r.interval = interval
	// We don't need to defer processing for the inbound requests.
	// As a result, we can create a non-buffered channel to store
	// requests
	r.requests = make(chan waiter)
	go r.poll()
	return &r
}

// RateLimiter manages the rate at which a request can be made
type RateLimiter struct {
	mux sync.RWMutex

	interval time.Duration
	requests chan waiter

	closed bool
}

func (r *RateLimiter) poll() {
	// Iterate through inbound requests
	for w := range r.requests {
		// Tell waiter to resume
		w.resume()
		// Sleep for delay interval
		time.Sleep(r.interval)
	}
}

// Acquire will request the next available window
func (r *RateLimiter) Acquire() {
	r.mux.RLock()
	defer r.mux.RUnlock()
	// Check to see if rate limiter has been closed
	if r.closed {
		// Rate limiter is closed, bail out
		return
	}

	// Create new waiter
	w := make(waiter)
	// Push waiter to queue
	r.requests <- w
	// Wait until resumed
	w.wait()
}

// Close will close an instance of rate limiter
func (r *RateLimiter) Close() (err error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	// Check to see if rate limiter has been closed
	if r.closed {
		// Rate limiter has already been closed, return error
		return errors.ErrIsClosed
	}

	// Close requests channel
	close(r.requests)
	// Set closed state to true
	r.closed = true
	return
}
