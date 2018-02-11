package easy_server

import(
	"sync"
	"runtime"
)

type worker struct{
	waitGroup * sync.WaitGroup
	dataFuncChan chan tcpDataFuncPacket
	udpDataFuncChan chan udpDataFuncPacket
}

func (w * worker) handleTcpPacket(workerNum int){
	defer w.waitGroup.Done()
	runtime.LockOSThread()
	for{
		select{
		case df := <-w.dataFuncChan:
			Logger.DebugLog("This is worker ",workerNum," handles the data ",df.bytes," for connection ",df.conn.conn.RemoteAddr())
			df.handlers.handleNoFirstPacket(df.conn,df.bytes)
		case udf := <-w.udpDataFuncChan:
		       	Logger.DebugLog("This is worker ",workerNum," handles the data ",udf.bytes," for udp connection ",udf.conn.addr)
			udf.handler(udf.conn,udf.bytes)
		}
	}
}
