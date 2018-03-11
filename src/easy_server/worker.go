package easy_server

import(
	"sync"
	// "fmt"
)

type worker struct{
	waitGroup * sync.WaitGroup
	dataFuncChan chan tcpDataFuncPacket
}

func (w * worker) handleTcpPacket(workerNum int){
	defer w.waitGroup.Done()
	for{
		select{
		case df := <-w.dataFuncChan:
			if df.handlers != nil{
				Logger.DebugLog("This is worker ",workerNum," handles the data ",df.bytes," for connection ",df.conn.conn.RemoteAddr())
				df.handlers.handleNoFirstPacket(df.conn,df.bytes)
			} else{
				return
			}
		}
	}
}
