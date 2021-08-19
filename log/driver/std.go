package driver

import (
	"log"
	"runtime/debug"
)

type Std struct {
	Log *log.Logger
	Name string
}

func (s *Std) Fatal(records ...interface{}) {
	s.Log.Fatalln(records)
}

func (s *Std) Err(records ...interface{}) {
	s.Log.Println(records)
}

func (s *Std) Warning(records ...interface{}) {
	s.Log.Println(records)
}

func (s *Std) Info(records ...interface{}) {
	s.Log.Println(records)
}

func (s *Std) Debug(record ...interface{}) {
	s.Log.Println(message("debug",s.Name,record))
}

func (s *Std) Trace(records ...interface{}) {
	s.Log.Println(records,string(debug.Stack()))
}

