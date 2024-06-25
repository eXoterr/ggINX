package logger

import "io"

type Config struct {
	outType       string
	writer        io.Writer
	level         string
	withTimestamp bool
}
