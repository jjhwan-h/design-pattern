package main

import (
	"context"
	"sync"
	"time"
)

type BroadCaster interface {
	Register(ctx context.Context, sub Subscriber, wg *sync.WaitGroup) chan Event
	UnRegister(ch chan Event)
	Publish(e Event)
}

type Event struct {
	Message string
	Time    time.Time
}

type broadCaster struct {
	Subs map[chan Event]struct{}

	mu sync.Mutex
}

func NewBroadCaster() *broadCaster {
	return &broadCaster{
		Subs: make(map[chan Event]struct{}),
	}
}

func (b *broadCaster) Register(ctx context.Context, sub Subscriber, wg *sync.WaitGroup) chan Event {
	ch := b.register()

	sub.Subscribe(ctx, ch, wg)
	go func() {
		<-ctx.Done()
		b.UnRegister(ch)
	}()

	return ch
}

func (b *broadCaster) register() chan Event {
	ch := make(chan Event, 1)

	b.mu.Lock()
	defer b.mu.Unlock()
	b.Subs[ch] = struct{}{}

	return ch
}

func (b *broadCaster) UnRegister(ch chan Event) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.Subs[ch]; ok {
		delete(b.Subs, ch)
		close(ch)
	}
}

func (b *broadCaster) Publish(e Event) {
	for ch := range b.Subs {
		ch <- e // blocks
	}
}

func (b *broadCaster) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	for ch := range b.Subs {
		close(ch)
		delete(b.Subs, ch)
	}
}
