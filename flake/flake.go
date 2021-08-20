package flake

import (
	"github.com/ccb1900/tinyweb/log"
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

func NextId() uint64  {
	id, err := sf.NextID()
	if err != nil {
		return 0
	}

	return id
}

func Init()  {
	local,err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
	//start := time.Now().Add(time.Duration(-1000))
	start := time.Date(1900,1,1,0,0,0,0,local)
	sf = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      start,
		//MachineID: func() (uint16, error) {
		//	return 1000,nil
		//},
		//CheckMachineID: func(u uint16) bool {
		//	return true
		//},
	})
}
