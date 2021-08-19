package driver

import (
	"fmt"
	"time"
)

func message(level string,name string,record ...interface{}) string  {
	t := time.Now().Format("2006-01-02 15-04-05")
	return fmt.Sprintf("[name=%s,level=%s,time=%s],%s",name,level,t,record)
}
