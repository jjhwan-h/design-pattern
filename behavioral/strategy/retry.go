package main

import (
	"net/http"
	"time"
)

type RetryStrategy interface {
	// NexBackoff는 attempt를 받아 다음 대기시간을 리턴
	NextBackoff(attempt int) time.Duration

	// ShouldRetry는 요청 결과를 보고 재시도할지를 결정
	ShouldRetry(resp *http.Response, err error) bool

	// MaxAttempts는 최대 시도 횟수를 리턴
	MaxAttempts() int
}
