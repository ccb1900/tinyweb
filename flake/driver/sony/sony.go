package sony

import (
	"github.com/ccb1900/tinyweb/config"
	"github.com/ccb1900/tinyweb/flake/contract"
	"github.com/ccb1900/tinyweb/log"
	"github.com/sony/sonyflake"
	"time"
)

type Sony struct {
	Sf *sonyflake.Sonyflake
}

func (s *Sony) NextId() uint64 {
	id, err := s.Sf.NextID()
	if err != nil {
		return 0
	}

	return id
}

func New() contract.IFlake  {
	s := new(Sony)
	local,err := time.LoadLocation(config.Get("app.timezone"))
	if err != nil {
		log.Fatal(err)
	}
	//start := time.Now().Add(time.Duration(-1000))
	start := time.Date(1900,1,1,0,0,0,0,local)
	s.Sf = sonyflake.NewSonyflake(sonyflake.Settings{
		StartTime:      start,
		MachineID: func() (uint16, error) {
			return 1000,nil
		},
		CheckMachineID: func(u uint16) bool {
			return true
		},
	})

	return s
}
