package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync"
	"time"
)

type AsyncHandler struct {
	handler slog.Handler
	ch      chan []slog.Attr
	done    chan struct{}
	wg      sync.WaitGroup
}

func NewAsyncHandler(wrapped slog.Handler, bufferSize int) *AsyncHandler {
	h := &AsyncHandler{
		handler: wrapped,
		ch:      make(chan []slog.Attr, bufferSize),
		done:    make(chan struct{}),
	}

	h.wg.Add(1)
	go h.process()

	return h
}

func (h *AsyncHandler) process() {
	defer h.wg.Done()

	for {
		select {
		case attrs := <-h.ch:
			r := slog.NewRecord(time.Now(), slog.LevelInfo, "", 0)
			r.AddAttrs(attrs...)
			_ = h.handler.Handle(context.Background(), r)
		case <-h.done:
			for len(h.ch) > 0 {
				attrs := <-h.ch
				r := slog.NewRecord(time.Now(), slog.LevelInfo, "", 0)
				r.AddAttrs(attrs...)
				_ = h.handler.Handle(context.Background(), r)
			}
			return
		}
	}
}

func (h *AsyncHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

func (h *AsyncHandler) Handle(_ context.Context, r slog.Record) error {
	attrs := make([]slog.Attr, 0, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})

	attrs = append(attrs,
		slog.String("time", r.Time.Format(time.RFC3339Nano)),
		slog.String("message", r.Message),
		slog.String("level", r.Level.String()),
	)

	select {
	case h.ch <- attrs:
	default:
		_, _ = io.WriteString(os.Stderr, "async log buffer full, dropping message: "+r.Message+"\n")
	}

	return nil
}

func (h *AsyncHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &AsyncHandler{
		handler: h.handler.WithAttrs(attrs),
		ch:      h.ch,
		done:    h.done,
		wg:      h.wg,
	}
}

func (h *AsyncHandler) WithGroup(name string) slog.Handler {
	return &AsyncHandler{
		handler: h.handler.WithGroup(name),
		ch:      h.ch,
		done:    h.done,
		wg:      h.wg,
	}
}

func (h *AsyncHandler) Close() {
	close(h.done)
	h.wg.Wait()
}

func NewAsyncLogger(w io.Writer, bufferSize int) *slog.Logger {
	baseHandler := slog.NewJSONHandler(w, &slog.HandlerOptions{})

	asyncHandler := NewAsyncHandler(baseHandler, bufferSize)

	return slog.New(asyncHandler)
}
