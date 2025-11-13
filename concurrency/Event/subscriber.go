package main

import (
	"context"
	"fmt"
	"sync"
)

type Subscriber interface {
	Subscribe(ctx context.Context, ch chan Event, wg *sync.WaitGroup)
	DoSomething(e Event)
}

type subscriber struct {
	name string
}

func NewSubscriber(name string) *subscriber {
	return &subscriber{name}
}

func (s *subscriber) DoSomething(e Event) {
	fmt.Printf("[%s]: %s\n", s.name, e.Message)
}

func (s *subscriber) Subscribe(ctx context.Context, ch chan Event, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for e := range ch {
			s.DoSomething(e)
		}
	}()
}
