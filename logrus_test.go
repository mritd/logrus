package logrus

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var fixedTime = time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

func newEntry(level logrus.Level, msg string, fields logrus.Fields) *logrus.Entry {
	return &logrus.Entry{
		Time:    fixedTime,
		Level:   level,
		Message: msg,
		Data:    fields,
		Buffer:  &bytes.Buffer{},
	}
}

func TestFormat_AllLevels(t *testing.T) {
	f := NewPlainFormatter()

	tests := []struct {
		level logrus.Level
		want  string
	}{
		{logrus.PanicLevel, "PANC"},
		{logrus.FatalLevel, "FATL"},
		{logrus.ErrorLevel, "ERRO"},
		{logrus.WarnLevel, "WARN"},
		{logrus.InfoLevel, "INFO"},
		{logrus.DebugLevel, "DEBG"},
		{logrus.TraceLevel, "TRAC"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			out, err := f.Format(newEntry(tt.level, "hello", nil))
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got := string(out)
			if !strings.Contains(got, tt.want) {
				t.Errorf("expected level %s in output, got: %s", tt.want, got)
			}
			if !strings.HasSuffix(got, "\n") {
				t.Error("output should end with newline")
			}
		})
	}
}

func TestFormat_NoFields(t *testing.T) {
	f := NewPlainFormatter()
	out, _ := f.Format(newEntry(logrus.InfoLevel, "test message", nil))
	want := "2025-01-01 12:00:00 INFO test message\n"
	if got := string(out); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFormat_WithFields(t *testing.T) {
	f := NewPlainFormatter()
	fields := logrus.Fields{"name": "alice", "role": "admin"}
	out, _ := f.Format(newEntry(logrus.InfoLevel, "login", fields))
	got := string(out)

	want := "2025-01-01 12:00:00 INFO login name=alice role=admin\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFormat_FieldsSorted(t *testing.T) {
	f := NewPlainFormatter()
	fields := logrus.Fields{"z": 1, "a": 2, "m": 3}
	out, _ := f.Format(newEntry(logrus.InfoLevel, "msg", fields))
	got := string(out)

	if !strings.Contains(got, "a=2 m=3 z=1") {
		t.Errorf("fields not sorted, got: %s", got)
	}
}

func TestFormat_FieldValueTypes(t *testing.T) {
	f := NewPlainFormatter()

	tests := []struct {
		name   string
		fields logrus.Fields
		want   string
	}{
		{"string", logrus.Fields{"k": "v"}, "k=v"},
		{"string_with_space", logrus.Fields{"k": "hello world"}, `k="hello world"`},
		{"string_with_equals", logrus.Fields{"k": "a=1"}, `k="a=1"`},
		{"string_with_quote", logrus.Fields{"k": `say "hi"`}, `k="say \"hi\""`},
		{"int", logrus.Fields{"k": 42}, "k=42"},
		{"int64", logrus.Fields{"k": int64(100)}, "k=100"},
		{"uint64", logrus.Fields{"k": uint64(200)}, "k=200"},
		{"float64", logrus.Fields{"k": 3.14}, "k=3.14"},
		{"bool_true", logrus.Fields{"k": true}, "k=true"},
		{"bool_false", logrus.Fields{"k": false}, "k=false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, _ := f.Format(newEntry(logrus.InfoLevel, "test", tt.fields))
			got := string(out)
			if !strings.Contains(got, tt.want) {
				t.Errorf("expected %q in output, got: %s", tt.want, got)
			}
		})
	}
}

func TestFormat_LevelOutOfRange(t *testing.T) {
	f := NewPlainFormatter()
	entry := newEntry(logrus.Level(99), "unknown level", nil)
	out, err := f.Format(entry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := string(out)
	if !strings.Contains(got, "UNKN") {
		t.Errorf("out-of-range level should show UNKN, got: %s", got)
	}
}

func TestFormat_NilLevelDesc(t *testing.T) {
	f := &PlainFormatter{TimestampFormat: defaultTimestampFormat}
	out, err := f.Format(newEntry(logrus.InfoLevel, "test", nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := string(out)
	if !strings.Contains(got, "UNKN") {
		t.Errorf("nil LevelDesc should show UNKN, got: %s", got)
	}
}

func TestFormat_EmptyTimestampFormat(t *testing.T) {
	f := &PlainFormatter{LevelDesc: defaultLevelDesc}
	out, _ := f.Format(newEntry(logrus.InfoLevel, "test", nil))
	got := string(out)
	if !strings.Contains(got, "2025-01-01 12:00:00") {
		t.Errorf("empty TimestampFormat should use default, got: %s", got)
	}
}

func TestFormat_EmptyFields(t *testing.T) {
	f := NewPlainFormatter()
	out, _ := f.Format(newEntry(logrus.InfoLevel, "test", logrus.Fields{}))
	got := string(out)

	want := "2025-01-01 12:00:00 INFO test\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestFormat_NilBuffer(t *testing.T) {
	f := NewPlainFormatter()
	entry := &logrus.Entry{
		Time:    fixedTime,
		Level:   logrus.InfoLevel,
		Message: "test",
		Data:    nil,
	}
	out, err := f.Format(entry)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := "2025-01-01 12:00:00 INFO test\n"
	if got := string(out); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNewPlainFormatter(t *testing.T) {
	f := NewPlainFormatter()
	if f.TimestampFormat != defaultTimestampFormat {
		t.Errorf("TimestampFormat = %q, want %q", f.TimestampFormat, defaultTimestampFormat)
	}
	if len(f.LevelDesc) != 7 {
		t.Errorf("LevelDesc length = %d, want 7", len(f.LevelDesc))
	}
}

func BenchmarkFormat_NoFields(b *testing.B) {
	f := NewPlainFormatter()
	entry := newEntry(logrus.InfoLevel, "benchmark message", nil)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		entry.Buffer.Reset()
		_, _ = f.Format(entry)
	}
}

func BenchmarkFormat_WithFields(b *testing.B) {
	f := NewPlainFormatter()
	fields := logrus.Fields{
		"user":   "alice",
		"action": "login",
		"ip":     "192.168.1.1",
		"port":   8080,
	}
	entry := newEntry(logrus.InfoLevel, "benchmark message", fields)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		entry.Buffer.Reset()
		_, _ = f.Format(entry)
	}
}
