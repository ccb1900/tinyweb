package common

import "time"

type Common struct {
	BootTime time.Time
}

var c *Common

func Init()  {
	c = &Common{BootTime: time.Now()}
}

func Get() *Common  {
	return c
}
