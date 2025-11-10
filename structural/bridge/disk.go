package main

import "context"

func NewDisk(storage Storage) *Disk {
	return &Disk{storage: storage}
}

type Disk struct {
	storage Storage
}

func (d *Disk) SetStorage(storage Storage) {
	d.storage = storage
}

func (d *Disk) PutData(ctx context.Context, data []byte) error {
	return d.storage.Put(ctx, data)
}
func (d *Disk) GetData(ctx context.Context) ([]byte, error) {
	return d.storage.Get(ctx)
}
