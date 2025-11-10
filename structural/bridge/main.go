package main

import (
	"context"
	"log"
)

func main() {
	data := []byte{
		12, 34,
	}

	local := &Local{}
	remote := &Remote{}

	disk := NewDisk(local)
	if err := disk.PutData(context.Background(), data); err != nil {
		log.Fatal(err)
	}
	if res, err := disk.GetData(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("[disk] local storage data: %b", res)
	}

	disk.SetStorage(remote)
	if err := disk.PutData(context.Background(), data); err != nil {
		log.Fatal(err)
	}
	if res, err := disk.GetData(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("[disk] remote storage data: %b", res)
	}

	card := NewCard(local)
	if err := card.PutData(context.Background(), data); err != nil {
		log.Fatal(err)
	}
	if res, err := card.GetData(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("[memory-card] local storage data: %b", res)
	}

	card.SetStorage(remote)
	if err := card.PutData(context.Background(), data); err != nil {
		log.Fatal(err)
	}
	if res, err := card.GetData(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("[memory-card] remote storage data: %b", res)
	}
}
