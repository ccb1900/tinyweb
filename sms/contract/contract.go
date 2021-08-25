package contract

type ISms interface {
	Send(mobile string,content string)
}
