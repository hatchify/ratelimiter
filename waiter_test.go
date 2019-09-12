package ratelimiter

import (
	"testing"
	"time"
)

func TestWaiter_resume(t *testing.T) {
	w := make(waiter)
	go w.resume()
	w.wait()
}

func TestWaiter_wait(t *testing.T) {
	w := make(waiter)
	go func() {
		w.wait()
		t.Fatal("Wait completed when it shouldn't have")
	}()

	<-time.NewTimer(time.Millisecond * 200).C
}
