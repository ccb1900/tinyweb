package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var v *viper.Viper

func LoadConfig(path string)  {
	v = viper.New()
	v.SetConfigFile(path)
	beforeLoad(v)
	err := v.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		log.Fatalln(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}

//
func beforeLoad(v *viper.Viper)  {
	v.SetDefault("app.timezone","Asia/Shanghai")
}


func Get(key string) string  {
	return v.GetString(key)
}

func GetInt(key string) int  {
	return v.GetInt(key)
}
var lock sync.Mutex
// Set 锁的粒度太大
func Set(key string,value string)  {
	lock.Lock()
	v.Set(key,value)
	lock.Unlock()
}
