package port

type Logger interface {
	Info(msg string, fields ...interface{})
	Error(msg string, fields ...interface{})
	Warn(msg string, fields ...interface{})
}
