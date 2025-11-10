package main

import (
	"context"
	"fmt"
	"sync"
)

type Local struct {
	storage []byte

	mu sync.RWMutex
}

func (l *Local) Put(ctx context.Context, data []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	fmt.Println("Put Local Storage")
	l.storage = data

	return nil
}
func (l *Local) Get(ctx context.Context) ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	fmt.Println("Get Local Storage")
	return l.storage, nil
}
