package logger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Config holds logger setup options.
type Config struct {
	Level      slog.Leveler
	Writer     io.Writer
	AddSource  bool
	Color      bool
	Format     Format
	SourceRoot string

	colorSet bool
}

// Format indicates output format.
type Format string

const (
	FormatText Format = "text"
	FormatJSON Format = "json"
)

// Option customizes logger Config.
type Option func(*Config)

// WithLevel sets the log level.
func WithLevel(level slog.Leveler) Option {
	return func(c *Config) {
		c.Level = level
	}
}

// WithWriter sets the output writer.
func WithWriter(w io.Writer) Option {
	return func(c *Config) {
		if w != nil {
			c.Writer = w
		}
	}
}

// WithFormat sets output format (text/json).
func WithFormat(format Format) Option {
	return func(c *Config) {
		if format != "" {
			c.Format = format
		}
	}
}

// WithColor explicitly enables or disables colorized output.
func WithColor(enabled bool) Option {
	return func(c *Config) {
		c.Color = enabled
		c.colorSet = true
	}
}

// WithSource controls whether to emit source location.
func WithSource(enabled bool) Option {
	return func(c *Config) {
		c.AddSource = enabled
	}
}

// Init builds a slog Logger with colorized console output and source location,
// sets it as the default logger, and returns it.
func Init(opts ...Option) *slog.Logger {
	cfg := &Config{
		Level:      slog.LevelInfo,
		Writer:     os.Stdout,
		AddSource:  true,
		Format:     FormatText,
		SourceRoot: defaultSourceRoot(),
	}
	applyEnv(cfg)
	for _, opt := range opts {
		opt(cfg)
	}

	if !cfg.colorSet {
		cfg.Color = shouldUseColor(cfg.Writer)
	}

	baseHandler := newBaseHandler(cfg)
	handler := &contextHandler{Handler: baseHandler}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}

func newBaseHandler(cfg *Config) slog.Handler {
	options := &slog.HandlerOptions{
		Level:     cfg.Level,
		AddSource: cfg.AddSource,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			return replaceAttr(cfg.Color, cfg.SourceRoot, a)
		},
	}

	if cfg.Format == FormatJSON {
		cfg.Color = false
		return slog.NewJSONHandler(cfg.Writer, options)
	}

	return &consoleHandler{
		w:           cfg.Writer,
		level:       cfg.Level,
		addSource:   cfg.AddSource,
		replaceAttr: options.ReplaceAttr,
		color:       cfg.Color,
		sourceRoot:  cfg.SourceRoot,
	}
}

func replaceAttr(enableColor bool, sourceRoot string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.LevelKey:
		level, ok := valueToLevel(a.Value)
		if !ok {
			return a
		}
		upper := level.String()
		if enableColor {
			upper = colorize(level, upper)
		}
		a.Value = slog.StringValue(upper)
	case slog.TimeKey:
		if t, ok := valueToTime(a.Value); ok {
			a.Value = slog.StringValue(t.Local().Format("2006-01-02 15:04:05.000"))
		}
	case slog.SourceKey:
		if src, ok := a.Value.Any().(slog.Source); ok {
			file := trimSourcePath(sourceRoot, src.File)
			a.Value = slog.StringValue(fmt.Sprintf("%s:%d", file, src.Line))
		}
	}
	return a
}

func valueToLevel(v slog.Value) (slog.Level, bool) {
	switch v.Kind() {
	case slog.KindInt64:
		return slog.Level(v.Int64()), true
	case slog.KindString:
		return parseLevelValue(v.String())
	default:
		if lv, ok := v.Any().(slog.Level); ok {
			return lv, true
		}
		return slog.LevelInfo, false
	}
}

func parseLevelValue(s string) (slog.Level, bool) {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug, true
	case "info":
		return slog.LevelInfo, true
	case "warn", "warning":
		return slog.LevelWarn, true
	case "error", "err":
		return slog.LevelError, true
	default:
		return slog.LevelInfo, false
	}
}

func valueToTime(v slog.Value) (time.Time, bool) {
	switch v.Kind() {
	case slog.KindTime:
		return v.Time(), true
	default:
		if t, ok := v.Any().(time.Time); ok {
			return t, true
		}
		return time.Time{}, false
	}
}

type ctxKey string

const requestIDKey ctxKey = "request_id"

// WithRequestID stores request id into context for downstream logging.
func WithRequestID(ctx context.Context, id string) context.Context {
	if id == "" {
		return ctx
	}
	return context.WithValue(ctx, requestIDKey, id)
}

// RequestIDFromContext fetches request id.
func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Value(requestIDKey).(string); ok {
		return v
	}
	return ""
}

type contextHandler struct {
	slog.Handler
}

func (h *contextHandler) Handle(ctx context.Context, r slog.Record) error {
	if reqID := RequestIDFromContext(ctx); reqID != "" {
		r.AddAttrs(slog.String("request_id", reqID))
	}
	return h.Handler.Handle(ctx, r)
}

type consoleHandler struct {
	w           io.Writer
	level       slog.Leveler
	addSource   bool
	replaceAttr func([]string, slog.Attr) slog.Attr
	attrs       []slog.Attr
	groups      []string
	color       bool
	sourceRoot  string
}

func (h *consoleHandler) Enabled(_ context.Context, level slog.Level) bool {
	min := slog.LevelInfo
	if h.level != nil {
		min = h.level.Level()
	}
	return level >= min
}

func (h *consoleHandler) Handle(_ context.Context, r slog.Record) error {
	if !h.Enabled(nil, r.Level) {
		return nil
	}

	var buf bytes.Buffer
	ts := r.Time
	if ts.IsZero() {
		ts = time.Now()
	}
	timeStr := ts.Local().Format("2006-01-02 15:04:05.000")
	if h.color {
		timeStr = colorCyan + timeStr + colorReset
	}
	buf.WriteString(timeStr)
	buf.WriteByte(' ')

	lvl := strings.ToUpper(r.Level.String())
	if h.color {
		lvl = colorize(r.Level, lvl)
	}
	buf.WriteByte('[')
	buf.WriteString(lvl)
	buf.WriteString("] ")

	if h.addSource {
		src := sourceFromRecord(r)
		if src.File != "" {
			path := trimSourcePath(h.sourceRoot, src.File)
			if h.color {
				path = colorBlue + path + colorReset
			}
			buf.WriteString(path)
			buf.WriteByte(':')
			lineStr := fmt.Sprintf("%d", src.Line)
			if h.color {
				lineStr = colorBlue + lineStr + colorReset
			}
			buf.WriteString(lineStr)
			buf.WriteByte(' ')
		}
	}

	msg := r.Message
	if h.color {
		msg = colorize(r.Level, msg)
	}
	buf.WriteString(msg)

	attrs := make([]slog.Attr, 0, len(h.attrs)+r.NumAttrs())
	for _, a := range h.attrs {
		attrs = appendAttr(attrs, h.groups, h.replaceAttr, a)
	}
	r.Attrs(func(a slog.Attr) bool {
		attrs = appendAttr(attrs, h.groups, h.replaceAttr, a)
		return true
	})

	for _, a := range attrs {
		buf.WriteByte(' ')
		buf.WriteString(a.Key)
		buf.WriteByte('=')
		buf.WriteString(formatValue(a.Value))
	}

	buf.WriteByte('\n')
	_, err := h.w.Write(buf.Bytes())
	return err
}

func (h *consoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	merged := make([]slog.Attr, 0, len(h.attrs)+len(attrs))
	merged = append(merged, h.attrs...)
	merged = append(merged, attrs...)
	return &consoleHandler{
		w:           h.w,
		level:       h.level,
		addSource:   h.addSource,
		replaceAttr: h.replaceAttr,
		attrs:       merged,
		groups:      h.groups,
		color:       h.color,
		sourceRoot:  h.sourceRoot,
	}
}

func (h *consoleHandler) WithGroup(name string) slog.Handler {
	if name == "" {
		return h
	}
	groups := append([]string{}, h.groups...)
	groups = append(groups, name)
	return &consoleHandler{
		w:           h.w,
		level:       h.level,
		addSource:   h.addSource,
		replaceAttr: h.replaceAttr,
		attrs:       h.attrs,
		groups:      groups,
		color:       h.color,
		sourceRoot:  h.sourceRoot,
	}
}

func shouldUseColor(w io.Writer) bool {
	file, ok := w.(*os.File)
	if !ok {
		return false
	}
	info, err := file.Stat()
	if err != nil {
		return false
	}
	if (info.Mode() & os.ModeCharDevice) == 0 {
		return false
	}
	if term := os.Getenv("TERM"); term == "" || term == "dumb" {
		return false
	}
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	return true
}

func trimSourcePath(root, file string) string {
	if root != "" {
		if rel, err := filepath.Rel(root, file); err == nil && !strings.HasPrefix(rel, "..") {
			return rel
		}
	}
	return file
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
)

func colorize(level slog.Level, text string) string {
	switch {
	case level >= slog.LevelError:
		return colorRed + text + colorReset
	case level >= slog.LevelWarn:
		return colorYellow + text + colorReset
	case level >= slog.LevelInfo:
		return colorGreen + text + colorReset
	default:
		return colorBlue + text + colorReset
	}
}

func appendAttr(dst []slog.Attr, groups []string, replacer func([]string, slog.Attr) slog.Attr, a slog.Attr) []slog.Attr {
	if len(groups) > 0 {
		keyParts := append(append([]string{}, groups...), a.Key)
		a.Key = strings.Join(keyParts, ".")
	}
	if replacer != nil {
		a = replacer(groups, a)
	}
	if a.Equal(slog.Attr{}) {
		return dst
	}
	return append(dst, a)
}

func formatValue(v slog.Value) string {
	switch v.Kind() {
	case slog.KindString:
		return v.String()
	case slog.KindBool:
		return strconv.FormatBool(v.Bool())
	case slog.KindInt64:
		return fmt.Sprint(v.Int64())
	case slog.KindFloat64:
		return strconv.FormatFloat(v.Float64(), 'f', -1, 64)
	case slog.KindDuration:
		return v.Duration().String()
	case slog.KindTime:
		return v.Time().Local().Format("2006-01-02 15:04:05.000")
	default:
		return fmt.Sprint(v.Any())
	}
}

func sourceFromRecord(r slog.Record) slog.Source {
	if src := r.Source(); src != nil && src.File != "" {
		return *src
	}
	if pc := r.PC; pc != 0 {
		return pcToSource(pc)
	}
	if pc, file, line, ok := runtime.Caller(4); ok {
		return slog.Source{Function: runtime.FuncForPC(pc).Name(), File: file, Line: line}
	}
	return slog.Source{}
}

func defaultSourceRoot() string {
	wd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return wd
}

func pcToSource(pc uintptr) slog.Source {
	frames := runtime.CallersFrames([]uintptr{pc})
	frame, _ := frames.Next()
	return slog.Source{Function: frame.Function, File: frame.File, Line: frame.Line}
}

func applyEnv(cfg *Config) {
	if v := strings.TrimSpace(os.Getenv("LOG_LEVEL")); v != "" {
		if parsed, ok := parseLevel(v); ok {
			cfg.Level = parsed
		}
	}

	if v := strings.TrimSpace(os.Getenv("LOG_FORMAT")); v != "" {
		lower := strings.ToLower(v)
		if lower == string(FormatJSON) {
			cfg.Format = FormatJSON
		} else if lower == string(FormatText) {
			cfg.Format = FormatText
		}
	}

	if v := strings.TrimSpace(os.Getenv("LOG_SOURCE_ROOT")); v != "" {
		cfg.SourceRoot = v
	}

	if v := strings.TrimSpace(os.Getenv("LOG_SOURCE")); v != "" {
		cfg.AddSource = parseBool(v, cfg.AddSource)
	}

	if v := strings.TrimSpace(os.Getenv("LOG_COLOR")); v != "" {
		cfg.Color = parseBool(v, cfg.Color)
		cfg.colorSet = true
	}
}

func parseLevel(v string) (slog.Level, bool) {
	switch strings.ToLower(v) {
	case "debug":
		return slog.LevelDebug, true
	case "info":
		return slog.LevelInfo, true
	case "warn", "warning":
		return slog.LevelWarn, true
	case "error", "err":
		return slog.LevelError, true
	default:
		return slog.LevelInfo, false
	}
}

func parseBool(v string, fallback bool) bool {
	switch strings.ToLower(v) {
	case "1", "true", "yes", "y", "on":
		return true
	case "0", "false", "no", "n", "off":
		return false
	default:
		return fallback
	}
}
