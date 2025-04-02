package logger

// Logger - интерфейс для логирования
type Logger interface {
	Info(msg string, fields map[string]interface{})
	Infof(msg string, fields map[string]interface{}, args ...interface{})
	Error(msg string, fields map[string]interface{})
	Errorf(msg string, fields map[string]interface{}, args ...interface{})
	Warn(msg string, fields map[string]interface{})
	Warnf(msg string, fields map[string]interface{}, args ...interface{})
	Debug(msg string, fields map[string]interface{})
	Debugf(msg string, fields map[string]interface{}, args ...interface{})
}
