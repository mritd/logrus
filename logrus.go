// Package logrus provides a plain-text formatter for sirupsen/logrus.
package logrus

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const defaultTimestampFormat = "2006-01-02 15:04:05"

var defaultLevelDesc = []string{"PANC", "FATL", "ERRO", "WARN", "INFO", "DEBG", "TRAC"}

// PlainFormatter formats logrus entries as plain text.
type PlainFormatter struct {
	TimestampFormat string
	LevelDesc       []string
}

// NewPlainFormatter returns a PlainFormatter with default settings.
func NewPlainFormatter() *PlainFormatter {
	return &PlainFormatter{
		TimestampFormat: defaultTimestampFormat,
		LevelDesc:       defaultLevelDesc,
	}
}

func (f *PlainFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	levelText := "UNKN"
	if f.LevelDesc != nil && int(entry.Level) < len(f.LevelDesc) {
		levelText = f.LevelDesc[entry.Level]
	}

	buf := entry.Buffer
	if buf == nil {
		buf = &bytes.Buffer{}
	}

	buf.WriteString(entry.Time.Format(timestampFormat))
	buf.WriteByte(' ')
	buf.WriteString(levelText)
	buf.WriteByte(' ')
	buf.WriteString(entry.Message)

	writeFields(buf, entry.Data)

	buf.WriteByte('\n')
	return buf.Bytes(), nil
}

func writeFields(buf *bytes.Buffer, data logrus.Fields) {
	if len(data) == 0 {
		return
	}

	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		buf.WriteByte(' ')
		buf.WriteString(k)
		buf.WriteByte('=')
		writeValue(buf, data[k])
	}
}

func writeValue(buf *bytes.Buffer, val interface{}) {
	switch v := val.(type) {
	case string:
		if needsQuote(v) {
			buf.WriteString(strconv.Quote(v))
		} else {
			buf.WriteString(v)
		}
	case int:
		buf.WriteString(strconv.Itoa(v))
	case int64:
		buf.WriteString(strconv.FormatInt(v, 10))
	case uint64:
		buf.WriteString(strconv.FormatUint(v, 10))
	case float64:
		buf.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
	case bool:
		buf.WriteString(strconv.FormatBool(v))
	default:
		s := fmt.Sprintf("%v", v)
		if needsQuote(s) {
			buf.WriteString(strconv.Quote(s))
		} else {
			buf.WriteString(s)
		}
	}
}

func needsQuote(s string) bool {
	return strings.ContainsAny(s, " \t=\"")
}

// SetDefault configures the global logrus logger with PlainFormatter.
func SetDefault() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(NewPlainFormatter())
}

func init() {
	SetDefault()
}
