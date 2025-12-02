package observer

import (
	"context"
	"time"
)

type command int

const (
	commanFlush command = iota
	commandFlushAndWait
	commandFlushDone
)

const (
	defaultTickerPeriod = 1 * time.Second
)

type handler[T any] struct {
	queue        *queue[T]
	fn           EventHandler[T]
	commandCh    chan command
	tickerUpdate chan time.Duration
	ticker       *time.Ticker
	tickerPeriod time.Duration
}

func newHandler[T any](queue *queue[T], fn EventHandler[T]) *handler[T] {
	return &handler[T]{
		queue:        queue,
		fn:           fn,
		commandCh:    make(chan command),
		tickerUpdate: make(chan time.Duration, 1),
		tickerPeriod: defaultTickerPeriod,
	}
}

func (h *handler[T]) withTick(period time.Duration) *handler[T] {
	h.tickerPeriod = period
	if h.ticker != nil {
		h.tickerUpdate <- period
	}
	return h
}

func (h *handler[T]) listen(ctx context.Context) {
	h.ticker = time.NewTicker(h.tickerPeriod)

	for {
		select {
		case <-h.ticker.C:
			go h.handle(ctx)
		case cmd, ok := <-h.commandCh:
			if !ok {
				return
			}

			h.handle(ctx)
			if cmd == commandFlushAndWait {
				h.ticker.Stop()
				close(h.commandCh)
			}
		case newPeriod := <-h.tickerUpdate:
			h.ticker.Stop()
			h.ticker = time.NewTicker(newPeriod)
		}
	}
}

func (h *handler[T]) handle(ctx context.Context) {
	h.fn(ctx, h.queue.All())
}

func (h *handler[T]) flush() {
	h.commandCh <- commanFlush
}

func (h *handler[T]) flushAndWait() {
	h.commandCh <- commandFlushAndWait
	<-h.commandCh
}
