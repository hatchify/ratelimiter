package ratelimiter

func newRequest(action func()) (r request) {
	// Set action
	r.action = action
	// Initialize waiter
	r.waiter = make(waiter)
	return
}

type request struct {
	// The action to be performed during the request
	action func()
	// Waiter co-ordinates the wait/resume actions for the request
	waiter waiter
}
