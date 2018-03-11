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
	tcpDataFuncPacketCh chan tcpDataFuncPacket
	connAlive bool
}

/*
create a new TcpConnection object
*/
func newTcpConnection(c net.Conn,cap int) *TcpConnection{
	t := &TcpConnection{
		tcpDataFuncPacketCh : make(chan tcpDataFuncPacket,cap),
		connAlive : true,
		conn : c,
    }

    return t
}

func (t * TcpConnection) SendData(bytes []byte){
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	if t.connAlive {
	   Logger.DebugLog("send data ",bytes," to conneciton ",t.conn.RemoteAddr())
		t.conn.Write(bytes)
	}
}

func (t * TcpConnection) Close(){
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	if t.connAlive {
		Logger.DebugLog("close conneciton ",t.conn.RemoteAddr())
		close(t.tcpDataFuncPacketCh)
		t.connAlive = false
	}
}
