package infra

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// --- Cosmetic Handler Interface Implementation ---

type cosmeticHandler struct {
	out      io.Writer
	opts     slog.HandlerOptions
	preAttrs []slog.Attr
	groups   []string
}

func (h *cosmeticHandler) prefix(key string) string {
	if len(h.groups) == 0 {
		return key
	}
	return strings.Join(h.groups, ".") + "." + key
}

func (h *cosmeticHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *cosmeticHandler) Handle(_ context.Context, r slog.Record) error {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "%s %-5s %s", r.Time.Format("15:04:05"), r.Level, r.Message)

	for _, a := range h.preAttrs {
		fmt.Fprintf(&buf, " %s=%v", h.prefix(a.Key), a.Value)
	}
	r.Attrs(func(a slog.Attr) bool {
		fmt.Fprintf(&buf, " %s=%v", h.prefix(a.Key), a.Value)
		return true
	})

	buf.WriteByte('\n')
	_, err := h.out.Write(buf.Bytes())
	return err
}

func (h *cosmeticHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	clone := *h
	clone.preAttrs = append(clone.preAttrs, attrs...)
	return &clone
}

func (h *cosmeticHandler) WithGroup(name string) slog.Handler {
	clone := *h
	clone.groups = append(clone.groups, name)
	return &clone
}

// -------------------------------------------------


// --- Multi Handler Interface Implementation ---
// The intent of this implementation is to be able to write nice
// logs to stdout while having the possibility to write struct json
// logs to a given file.

type multiHandler struct {
	handlers []slog.Handler
}

func (m *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.handlers {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m *multiHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, h := range m.handlers {
		if err := h.Handle(ctx, r.Clone()); err != nil {
			return err
		}
	}
	return nil
}

func (m *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		handlers[i] = h.WithAttrs(attrs)
	}
	return &multiHandler{handlers: handlers}
}

func (m *multiHandler) WithGroup(name string) slog.Handler {
	handlers := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		handlers[i] = h.WithGroup(name)
	}
	return &multiHandler{handlers: handlers}
}

// ----------------------------------------------


func SetupLogger(file *io.Writer){
	opts := slog.HandlerOptions{Level: slog.LevelInfo}
	handler := &cosmeticHandler{out: os.Stdout, opts: opts}
	
	var logger *slog.Logger
	if file != nil {
		logger = slog.New(&multiHandler{handlers: []slog.Handler{
			handler,
			slog.NewJSONHandler(*file, &opts),
		}})
	} else {
		logger = slog.New(handler)
	}

	slog.SetDefault(logger)
}
