package main

import (
	"net/http"
	"time"
)

type ConstantBackoff struct{}

func NewConstantBackoff() *ConstantBackoff {
	return &ConstantBackoff{}
}

func (c *ConstantBackoff) NextBackoff(attempt int) time.Duration {
	return time.Duration(time.Second)
}

func (c *ConstantBackoff) ShouldRetry(resp *http.Response, err error) bool {
	return true
}

func (c *ConstantBackoff) MaxAttempts() int {
	return 0
}
