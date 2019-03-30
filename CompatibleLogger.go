package zaplogger

import (
	"fmt"

	"go.uber.org/zap"
)

// ZapLogger is a logger which compatible to logrus/std zlog/prometheus.
// it implements Print() Println() Printf() Dbug() Debugln() Debugf() Info() Infoln() Infof() Warn() Warnln() Warnf()
// Error() Errorln() Errorf() Fatal() Fataln() Fatalf() Panic() Panicln() Panicf() With() WithField() WithFields()

type Logger struct {
	Log *zap.Logger
}

type ZapLogger = Logger

// NewProduction new log for production
func NewProduction() *Logger {
	log, _ := zap.NewProduction()
	return &ZapLogger{
		Log: log,
	}
}

// NewDevelopment new log for development
func NewDevelopment() *Logger {
	log, _ := zap.NewDevelopment()
	return &ZapLogger{
		Log: log,
	}
}

// InitZaoLogger initial
func InitZaoLogger(log *zap.Logger) *Logger {
	return &ZapLogger{
		log,
	}
}

// NewZapLogger return ZapLogger with caller field
func NewZapLogger() *Logger {
	return &ZapLogger{NewLogger().WithOptions(zap.AddCallerSkip(1))}
}

// Debug logs a message at level Debug on the ZapLogger.
func (l *Logger) Debug(args ...interface{}) {
	l.Log.Debug(fmt.Sprint(args...))
}

// Debugf logs a message at level Debug on the ZapLogger.
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.Log.Debug(fmt.Sprintf(template, args...))
}

// Info logs a message at level Info on the ZapLogger.
func (l *Logger) Info(args ...interface{}) {
	l.Log.Info(fmt.Sprint(args...))
}

// Infof logs a message at level Info on the ZapLogger.
func (l *Logger) Infof(template string, args ...interface{}) {

	l.Log.Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at level Warn on the ZapLogger.
func (l *Logger) Warn(args ...interface{}) {
	l.Log.Warn(fmt.Sprint(args...))
}

// Warnf logs a message at level Warn on the ZapLogger.
func (l *Logger) Warnf(template string, args ...interface{}) {

	l.Log.Warn(fmt.Sprintf(template, args...))
}

// Error logs a message at level Error on the ZapLogger.
func (l *Logger) Error(args ...interface{}) {
	l.Log.Error(fmt.Sprint(args...))
}

// Errorf logs a message at level Warn on the ZapLogger.
func (l *Logger) Errorf(template string, args ...interface{}) {

	l.Log.Error(fmt.Sprintf(template, args...))
}

// Fatal logs a message at level Fatal on the ZapLogger.
func (l *Logger) Fatal(args ...interface{}) {
	l.Log.Fatal(fmt.Sprint(args...))
}

// Fatalf logs a message at level Warn on the ZapLogger.
func (l *Logger) Fatalf(template string, args ...interface{}) {

	l.Log.Fatal(fmt.Sprintf(template, args...))
}

// Panic logs a message at level Painc on the ZapLogger.
func (l *Logger) Panic(args ...interface{}) {
	l.Log.Panic(fmt.Sprint(args...))
}

// Panicf  logs a message at level Warn on the ZapLogger.
func (l *Logger) Panicf(template string, args ...interface{}) {

	l.Log.Panic(fmt.Sprintf(template, args...))
}

// With return a logger with an extra field.
func (l *Logger) With(key string, value interface{}) *Logger {
	return &ZapLogger{l.Log.With(zap.Any(key, value))}
}

// Printf logs a message at level Info on the ZapLogger.
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Log.Info(fmt.Sprintf(format, args...))
}

// Print logs a message at level Info on the ZapLogger.
func (l *Logger) Print(args ...interface{}) {
	l.Log.Info(fmt.Sprint(args...))
}

// WithField return a logger with an extra field.
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &ZapLogger{l.Log.With(zap.Any(key, value))}
}

// WithFields return a logger with extra fields.
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	i := 0
	var clog *Logger
	for k, v := range fields {
		if i == 0 {
			clog = l.WithField(k, v)
		} else {
			clog = clog.WithField(k, v)
		}
		i++
	}
	return clog
}
