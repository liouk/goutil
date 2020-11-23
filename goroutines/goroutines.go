package goroutines

import "time"

// WaitTimeout dispatches a func of work in a goroutine, and waits for timeout for it to complete, or else returns false
func WaitTimeout(timeout time.Duration, work func()) (timedOut bool) {
	c := make(chan struct{})
	go func() {
		defer close(c)
		work()
	}()

	select {
	case <-c:
		timedOut = false
	case <-time.After(timeout):
		timedOut = true
	}

	return
}
