package contract

import "time"

type ICache interface {
	Get(key string) interface{}
	Set(key string,value interface{})
	SetWithExpire(key string,value interface{},expire time.Time)
	Delete(key string) error

}
