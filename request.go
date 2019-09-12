package ratelimiter

func newRequest(fn func()) (r request) {
	r.action = fn
	r.waiter = make(waiter)
	return
}

type request struct {
	action func()
	waiter waiter
}
