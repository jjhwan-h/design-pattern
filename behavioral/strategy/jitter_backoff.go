package main

import (
	"net/http"
	"time"
)

type JitterBackoff struct{}

func NewJitterBackoff() *JitterBackoff {
	return &JitterBackoff{}
}

func (j *JitterBackoff) NextBackoff(attempt int) time.Duration {
	return time.Duration(time.Second)
}

func (j *JitterBackoff) ShouldRetry(resp *http.Response, err error) bool {
	return true
}

func (j *JitterBackoff) MaxAttempts() int {
	return 0
}
