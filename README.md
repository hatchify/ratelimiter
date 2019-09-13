# RateLimiter
RateLimiter will limit the rate at which a request is made

## Usage
Usage examples of all exported methods are available below:

### New
```go
func ExampleNew() {
	// Initialize rate limiter
	r := New(time.Second)
	fmt.Printf("RateLimiter (%v) is ready to use!\n", r)
}
```

### RateLimiter.Acquire
```go
func ExampleRateLimiter_Acquire() {
	// Initialize rate limiter
	r := New(time.Second)
	// Acquire a new request
	r.Acquire()
}
```

### RateLimiter.Close 
```go
func ExampleRateLimiter_Close() {
	// Initialize rate limiter
	r := New(time.Second)

	// Do lots of great tasks here

	// Close rate limiter
	r.Close()
}
```
