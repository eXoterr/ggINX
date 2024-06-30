package logger

import (
	"bytes"
	"io"
	"testing"
)

const (
	LEVEL_INFO  = "info"
	LEVEL_WARN  = "warn"
	LEVEL_ERROR = "error"
	LEVEL_DEBUG = "debug"
)

func TestLoggerText(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})

	cfg := &Config{
		OutType:       "text",
		Writer:        buffer,
		Level:         LEVEL_INFO,
		WithTimestamp: false,
	}

	log := New()
	err := log.Setup(cfg)
	if err != nil {
		t.Errorf("unable to create logger: %s", err)
	}

	testingTable := []struct {
		outType  string
		msg      string
		level    string
		expected string
		withAttr Attribute
	}{
		{
			outType:  "text",
			msg:      "test debug message",
			level:    LEVEL_DEBUG,
			expected: "",
		},
		{
			outType:  "text",
			msg:      "test info message",
			level:    LEVEL_INFO,
			expected: "level=INFO msg=\"test info message\"\n",
		},
		{
			outType:  "text",
			msg:      "test info message",
			level:    LEVEL_INFO,
			expected: "level=INFO msg=\"test info message\" module=logger\n",
			withAttr: Attribute{Key: "module", Value: "logger"},
		},
	}

	for _, test := range testingTable {
		if test.withAttr.Key != "" {
			log = log.WithAttribute(test.withAttr)
		}

		switch test.level {
		case LEVEL_DEBUG:
			log.Debug(test.msg)
		case LEVEL_INFO:
			log.Info(test.msg)
		case LEVEL_WARN:
			log.Warn(test.msg)
		case LEVEL_ERROR:
			log.Error(test.msg)
		}

		result, err := io.ReadAll(buffer)
		if err != nil {
			t.Errorf("unable to read log result: %s", err)
		}

		actual := string(result)

		if actual != test.expected {
			t.Errorf("log result is incorrect: %s != %s", actual, test.expected)
		}
	}
}

func TestLoggerJSON(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{})

	cfg := &Config{
		OutType:       "json",
		Writer:        buffer,
		Level:         LEVEL_INFO,
		WithTimestamp: false,
	}

	log := New()
	err := log.Setup(cfg)
	if err != nil {
		t.Errorf("unable to create logger: %s", err)
	}

	testingTable := []struct {
		outType  string
		msg      string
		level    string
		expected string
		withAttr Attribute
	}{
		{
			outType:  "text",
			msg:      "test debug message",
			level:    LEVEL_DEBUG,
			expected: "",
		},
		{
			outType:  "text",
			msg:      "test info message",
			level:    LEVEL_INFO,
			expected: "{\"level\":\"INFO\",\"msg\":\"test info message\"}\n",
		},
		{
			outType:  "text",
			msg:      "test info message",
			level:    LEVEL_INFO,
			expected: "{\"level\":\"INFO\",\"msg\":\"test info message\",\"module\":\"logger\"}\n",
			withAttr: Attribute{Key: "module", Value: "logger"},
		},
	}

	for _, test := range testingTable {
		if test.withAttr.Key != "" {
			log = log.WithAttribute(test.withAttr)
		}

		switch test.level {
		case LEVEL_DEBUG:
			log.Debug(test.msg)
		case LEVEL_INFO:
			log.Info(test.msg)
		case LEVEL_WARN:
			log.Warn(test.msg)
		case LEVEL_ERROR:
			log.Error(test.msg)
		}

		result, err := io.ReadAll(buffer)
		if err != nil {
			t.Errorf("unable to read log result: %s", err)
		}

		actual := string(result)

		if actual != test.expected {
			t.Errorf("log result is incorrect: %s != %s", actual, test.expected)
		}
	}
}
