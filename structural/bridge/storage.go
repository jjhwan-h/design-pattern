package main

import "context"

// Implementor

type Storage interface {
	Put(ctx context.Context, data []byte) error
	Get(ctx context.Context) ([]byte, error)
}
