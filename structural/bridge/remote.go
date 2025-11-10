package main

import (
	"context"
	"fmt"
	"sync"
)

type Remote struct {
	storage []byte

	mu sync.RWMutex
}

func (r *Remote) Put(ctx context.Context, data []byte) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	fmt.Println("Put Remote Storage")
	r.storage = data

	return nil
}
func (r *Remote) Get(ctx context.Context) ([]byte, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	fmt.Println("Get Remote Storage")
	return r.storage, nil
}
