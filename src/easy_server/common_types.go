package easy_server

type SplitError int
const (
	NoSplitError SplitError = iota
	LessDataSplitError
	OtherSplitError
)

type tcpDataFuncPacket struct{
	conn * TcpConnection
	bytes []byte
	handlers * TcpDataHandlers
}

type TcpDataHandlers struct{
	splitPacket func([]byte) (int,SplitError)
	handleFirstPacket func(TcpConnectionOps,[]byte)
	handleNoFirstPacket func(TcpConnectionOps,[]byte)
	workerNum int
}


func NewTcpDataHandlers(s func([]byte) (int,SplitError),hf,ho func(TcpConnectionOps,[]byte),workernum int) * TcpDataHandlers{
     if hf ==nil || ho ==nil {
     	panic("The tcp packet handler must not be nil")
	return nil
     }
     t := &TcpDataHandlers{s,hf,ho,workernum}
     return t
}
