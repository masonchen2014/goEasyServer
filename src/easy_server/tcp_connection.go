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
	connAlive bool
}

/*
create a new TcpConnection object
*/
func newTcpConnection(c net.Conn) *TcpConnection{
	t := &TcpConnection{
		closeCh : make(chan struct{}),
		dataCh : make(chan []byte),
		connAlive : true,
		conn : c,
    }
    return t
}

func (t * TcpConnection) SendData(bytes []byte){
	t.rwMutex.RLock()
	defer t.rwMutex.RUnlock()
	if t.connAlive {
		t.dataCh <- bytes
	}
}

func (t * TcpConnection) Close(){
	t.rwMutex.Lock()
	defer t.rwMutex.Unlock()
	if t.connAlive {
		close(t.closeCh)
                close(t.dataCh)
		t.connAlive = false
	}
}
