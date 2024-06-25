package logger

type Attribute struct {
	Key   string
	Value interface{}
}

type Logger interface {
	Debug(msg string, any ...interface{})
	Error(msg string, any ...interface{})
	Info(msg string, any ...interface{})
	Warn(msg string, any ...interface{})

	Setup(config *Config) error

	WithAttribute(attribute Attribute) Logger
}
