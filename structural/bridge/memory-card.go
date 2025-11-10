package main

import "context"

func NewCard(storage Storage) *Card {
	return &Card{storage: storage}
}

type Card struct {
	storage Storage
}

func (c *Card) SetStorage(storage Storage) {
	c.storage = storage
}

func (c *Card) PutData(ctx context.Context, data []byte) error {
	return c.storage.Put(ctx, data)
}
func (c *Card) GetData(ctx context.Context) ([]byte, error) {
	return c.storage.Get(ctx)
}
