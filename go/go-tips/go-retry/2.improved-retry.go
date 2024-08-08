package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// set rand seed
func init(){
	rand.Seed(time.Now().UnixNano())
}

// DoRetry implements retry mechanism, stop retry when encounter StopRetry
// refererce: https://upgear.io/blog/simple-golang-retry-function/
func DoRetry(attempts int, sleep time.Duration, f func() error) error {
	if err := f(); err != nil {
		if s, ok := err.(stopRetryErr); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating a Thundering Herd
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2
			fmt.Printf("retry func error: %s. attemps #%d after %s.\n", err.Error(), attempts, sleep)
			time.Sleep(sleep)
			return DoRetry(attempts, 2*sleep, f)
		}
		return err
	}

	return nil
}

type stopRetryErr struct {
	error
}

func StopRetry(err error) error {
	return stopRetryErr{err}
}

func alwaysFailed() (string, error) {
	return "failed", errors.New("always failed")
}

func retryOnce() (string, error) {
	return "retry once", StopRetry(errors.New("retry once"))
}

func main() {
	var resp string
	var sendErr error
	_ = DoRetry(3, time.Second, func() error {
		resp, sendErr = alwaysFailed()
		return sendErr
	})
	fmt.Printf("resp: %s\n", resp)

	_ = DoRetry(3, time.Second, func() error {
		resp, sendErr = retryOnce()
		return sendErr
	})
	fmt.Printf("resp: %s", resp)
}

// t1 = t0 + 1/2 * [0, t0)
// t2 = 2*t1 + [0, t0)
// t3 = 2*t2 + [0, t2)
// reference: https://upgear.io/blog/simple-golang-retry-function/