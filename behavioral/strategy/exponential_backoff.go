package main

import (
	"net/http"
	"time"
)

type ExponentialBackoff struct {
}

func NewExponentialBackoff() *ExponentialBackoff {
	return &ExponentialBackoff{}
}

func (e *ExponentialBackoff) NextBackoff(attempt int) time.Duration {
	return time.Duration(time.Second)
}

func (e *ExponentialBackoff) ShouldRetry(resp *http.Response, err error) bool {
	return true
}

func (e *ExponentialBackoff) MaxAttempts() int {
	return 0
}
