package easy_server

import(
	"sync"
	"runtime"
)

type worker struct{
	waitGroup * sync.WaitGroup
	dataFuncChan chan tcpDataFuncPacket
}

func (w * worker) handleTcpPacket(workerNum int){
	defer w.waitGroup.Done()
	runtime.LockOSThread()
	for{
		select{
		case df := <-w.dataFuncChan:
			Logger.DebugLog("This is worker ",workerNum," handles the data ",df.bytes," for connection ",df.conn.conn.RemoteAddr())
			df.handlers.handleNoFirstPacket(df.conn,df.bytes)
		}
	}
}
