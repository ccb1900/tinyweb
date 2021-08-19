package driver

type Zap struct {
	
}

func (z *Zap) Fatal(records ...interface{}) {
	panic("implement me")
}

func (z *Zap) Err(records ...interface{}) {
	panic("implement me")
}

func (z *Zap) Warning(records ...interface{}) {
	panic("implement me")
}

func (z *Zap) Info(records ...interface{}) {
	panic("implement me")
}

func (z *Zap) Debug(records ...interface{}) {
	panic("implement me")
}

func (z *Zap) Trace(records ...interface{}) {
	panic("implement me")
}

