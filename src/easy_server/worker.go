package easy_server

import(
	"fmt"
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
			fmt.Printf("This is woker %d\n",workerNum)
			df.handlers.handleNoFirstPacket(df.ops,df.bytes)
		}
	}
}
