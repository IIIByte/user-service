package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"runtime"
)

// Log - реализация интерфейса Logger с использованием logrus
type Log struct {
	logger *logrus.Logger
}

// NewLogger создает новый экземпляр логгера с JSON-форматированием
func NewLogger(level logrus.Level) *Log {
	l := logrus.New()

	// Устанавливаем вывод в stdout
	l.SetOutput(os.Stdout)

	// Устанавливаем уровень логирования
	l.SetLevel(level)

	// Включаем информацию о вызове
	l.SetReportCaller(true)

	// Настраиваем JSON-форматирование
	l.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "15:04:05 02-01-2006", // Читаемый формат времени
		PrettyPrint:     false,                 // Компактный вывод (без отступов)
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := filepath.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	// TODO: Добавить возможность записи в файл

	return &Log{l}
}

// Info - реализует метод интерфейса Logger
func (l *Log) Info(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields).Info(msg)
	} else {
		l.Info(msg, fields)
	}
}

// Infof - реализует метод интерфейса Logger
func (l *Log) Infof(msg string, fields map[string]interface{}, args ...interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields).Info(msg, args)
	} else {
		l.Infof(msg, fields, args)
	}
}

// Error - реализует метод интерфейса Logger
func (l *Log) Error(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields).Error(msg)
	} else {
		l.Error(msg, fields)
	}
}

// Errorf - реализует метод интерфейса Logger
func (l *Log) Errorf(msg string, fields map[string]interface{}, args ...interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields).Errorf(msg, args)
	} else {
		l.Errorf(msg, fields, args)
	}
}

// Warn - реализует метод интерфейса Logger
func (l *Log) Warn(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields).Warn(msg)
	} else {
		l.Warn(msg, fields)
	}
}

// Warnf - реализует метод интерфейса Logger
func (l *Log) Warnf(msg string, fields map[string]interface{}, args ...interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields).Warnf(msg, args)
	} else {
		l.Warnf(msg, fields, args)
	}
}

// Debug - реализует метод интерфейса Logger
func (l *Log) Debug(msg string, fields map[string]interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields).Debug(msg)
	} else {
		l.Debug(msg, fields)
	}
}

// Debugf - реализует метод интерфейса Logger
func (l *Log) Debugf(msg string, fields map[string]interface{}, args ...interface{}) {
	if len(fields) > 0 {
		l.logger.WithFields(fields).Debugf(msg)
	} else {
		l.Debugf(msg, fields, args)
	}
}
