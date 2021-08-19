package log

import "github.com/ccb1900/tinyweb/config"

type ILog interface {
	Fatal(records ...interface{})
	Err(records ...interface{})
	Warning(records ...interface{})
	Info(records ...interface{})
	Debug(records ...interface{})
	Trace(records ...interface{})
}
var defaultLog ILog

func Fatal()  {
	defaultLog.Fatal()
}

func Error()  {
	defaultLog.Err()
}

func Warning()  {
	defaultLog.Warning()
}

func Info()  {
	defaultLog.Info()
}

func Debug(record ...interface{})  {
	defaultLog.Debug(record)
}
func Trace()  {
	defaultLog.Trace()
}

func Driver(name string)  {
	defaultLog = Factory(name)
}

func Init()  {
	Driver(config.Get("log.driver"))
}
