package ratelimiter

type waiter chan struct{}

func (w waiter) wait() {
	// Wait for empty struct to be pushed to waiter
	<-w
}

func (w waiter) resume() {
	// Push empty struct to waiter
	w <- struct{}{}
}
