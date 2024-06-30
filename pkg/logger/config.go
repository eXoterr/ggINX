package logger

import "io"

type Config struct {
	OutType       string
	Writer        io.Writer
	Level         string
	WithTimestamp bool
}
