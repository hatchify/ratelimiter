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
	r.requests = make(chan *request)
	go r.poll()
	return &r
}

// RateLimiter manages the rate at which a request can be made
type RateLimiter struct {
	mux sync.RWMutex

	interval time.Duration
	requests chan *request

	closed bool
}

func (r *RateLimiter) poll() {
	// Iterate through inbound requests
	for req := range r.requests {
		// Perform action
		req.action()
		// Action has completed, tell waiter to resume
		req.waiter.resume()
		// Sleep for delay interval
		time.Sleep(r.interval)
	}
}

// Acquire will request the next available window
func (r *RateLimiter) Acquire(fn func()) {
	r.mux.RLock()
	defer r.mux.RUnlock()
	// Check to see if rate limiter has been closed
	if r.closed {
		// Rate limiter is closed, bail out
		return
	}

	// Create a new request
	req := newRequest(fn)
	// Push request to queue
	r.requests <- &req
	// Wait until resumed
	req.waiter.wait()
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
