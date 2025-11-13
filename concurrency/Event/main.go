package main

import (
	"context"
	"sync"
	"time"
)

func main() {
	b := NewBroadCaster()

	sub1 := NewSubscriber("sub1")
	sub2 := NewSubscriber("sub2")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	ch1 := b.Register(ctx, sub1, &wg)
	b.Register(ctx, sub2, &wg)

	b.Publish(Event{Message: "hello", Time: time.Now()})
	b.Publish(Event{Message: "world", Time: time.Now()})

	b.UnRegister(ch1)
	b.Publish(Event{Message: "hello again", Time: time.Now()})

	cancel()
	wg.Wait()

	b.Close()
}
