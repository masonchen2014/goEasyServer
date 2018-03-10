package easy_server

import(
	"net"
	"sync"
)
/*
interface for tcp connection operations
*/
type TcpConnectionOps interface{
      SendData(bytes []byte)
      Close()
}

/*
type TcpConnection and its functions
*/
type TcpConnection struct{
	rwMutex sync.RWMutex
	conn net.Conn
	closeCh chan struct{}
	dataCh chan []byte
	tcpDataFuncPacketCh chan tcpDataFuncPacket
	connAlive bool
	// w worker
	// r receiver
}

/*
create a new TcpConnection object
*/
func newTcpConnection(c net.Conn,cap int,waitGroup *sync.WaitGroup) *TcpConnection{
	t := &TcpConnection{
		closeCh : make(chan struct{}),
		dataCh : make(chan []byte),
		tcpDataFuncPacketCh : make(chan tcpDataFuncPacket,cap),
		connAlive : true,
		conn : c,
    }
	// t.w = worker{waitGroup,t.tcpDataFuncPacketCh}
	// t.r = receiver{waitGroup,t.tcpDataFuncPacketCh}

    return t
}

func (t * TcpConnection) SendData(bytes []byte){
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	if t.connAlive {
	   Logger.DebugLog("send data ",bytes," to conneciton ",t.conn.RemoteAddr())
	   t.dataCh <- bytes
	}
}

func (t * TcpConnection) Close(){
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	if t.connAlive {
		Logger.DebugLog("close conneciton ",t.conn.RemoteAddr())
		close(t.closeCh)
        close(t.dataCh)
		close(t.tcpDataFuncPacketCh)
		t.connAlive = false
	}
}
