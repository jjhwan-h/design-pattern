package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// *http.Client와 동일한 최소 인터페이스
type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

// Strategy를 주입받아 재시도처리하는 래퍼
type Client struct {
	Doer     Doer
	Strategy RetryStrategy
}

func (c *Client) DoWithRetry(ctx context.Context, req *http.Request) (*http.Response, error) {
	if c.Doer == nil || c.Strategy == nil {
		return nil, errors.New("nil Doer or Strategy")
	}

	var lastErr error
	var resp *http.Response

	for attempt := 0; attempt < c.Strategy.MaxAttempts(); attempt++ {
		// 요청에 컨텍스트 주입 (타임아웃/캔슬 전파)
		req = req.Clone(ctx)

		resp, lastErr = c.Doer.Do(req)

		// 마지막 시도면 탈출
		if attempt == c.Strategy.MaxAttempts()-1 {
			break
		}

		if !c.Strategy.ShouldRetry(resp, lastErr) {
			break
		}

		// backoff 후 다음 시도
		wait := c.Strategy.NextBackoff(attempt)
		t := time.NewTimer(wait)
		select {
		case <-t.C:
			// 계속
		case <-ctx.Done():
			t.Stop()
			return nil, fmt.Errorf("retry canceled: %w", ctx.Err())
		}

		// 응답 바디가 있으면 닫아 줌(다음 시도 전 자원 해제)
		if resp != nil && resp.Body != nil {
			_ = resp.Body.Close()
		}
	}

	return resp, lastErr
}

func (c *Client) SetRetryStrategy(strategy RetryStrategy) {
	c.Strategy = strategy
}
