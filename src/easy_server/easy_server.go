package  easy_server
import(
	"net"
	"runtime"
	"sync"
	"time"
)

/*
type EasyServer and its functions
*/
type EasyServer struct{
	waitGroup * sync.WaitGroup
	num_of_listeners int
}

/*
create EasyServer object
*/
func NewServer() *EasyServer {
	s := &EasyServer{
		waitGroup: &sync.WaitGroup{},
    }
    return s
}

/*
listen on the specified port using tcp protocol
*/
func (server *EasyServer) AddTcpListener(port string,h * TcpDataHandlers){
 	server.waitGroup.Add(1)
	server.num_of_listeners++
	go server.addTcpListener(port,h)
	Logger.SysLog("Add Tcp listenner on port : "+port)
}

/*
print the internal information of EasyServer
*/
func (server * EasyServer) PrintServerInfo(){
	for{
		time.Sleep(2*time.Second)
		Logger.SysLog("number of listeners : ",server.num_of_listeners)
		Logger.SysLog("number of goroutines : ",runtime.NumGoroutine())
	}

}

/*
wait all the works done
*/
func (server * EasyServer) Stop(){
	Logger.SysLog("Wait for all the jobs done...")
	server.waitGroup.Wait()
}


func (server *EasyServer) addTcpListener(port string,h * TcpDataHandlers) {
	defer server.waitGroup.Done()
	ln, err := net.Listen("tcp",port)
	defer ln.Close()
	if err != nil {
        Logger.ErrorLog(err)
		panic("TCP can't listen on port "+port)
	}

    for {
        conn, err := ln.Accept()
        if err != nil {
            Logger.WarnLog(err)
            continue
        }

		server.waitGroup.Add(1)
        go server.handleConnection(conn,h)
    }

}

func (server *EasyServer) handleConnection(conn net.Conn,h * TcpDataHandlers){
	Logger.DebugLog("Accept a connection from ",conn.RemoteAddr())
	defer server.waitGroup.Done()

	t := newTcpConnection(conn,h.workerNum*2)
	Logger.DebugLog("create ",h.workerNum," workers for connection from ",conn.RemoteAddr())

	w := worker{server.waitGroup,t.tcpDataFuncPacketCh}
	for i:=0;i<h.workerNum;i++{
		server.waitGroup.Add(1)
		go w.handleTcpPacket(i)
	}

	r := receiver{t.tcpDataFuncPacketCh}
    r.splitPacket(t,h)
}

/*
listen on the specified port using udp protocol
*/
func (server *EasyServer) AddUdpListener(port string,h func(UdpPacketOps,[]byte)){
 	server.waitGroup.Add(1)
	server.num_of_listeners++
	go server.addUdpListener(port,h)
	Logger.SysLog("Add udp listener on port ",port)
}


func (server *EasyServer) addUdpListener(port string,h func(UdpPacketOps,[]byte)){
    // 创建监听
    pConn, err := net.ListenPacket("udp",port)
    defer pConn.Close()
    if err != nil {
        panic(err)
        return
    }

    data := make([]byte,65535)

	u := UdpPacketHandler{server.waitGroup}

    for {
        // 读取数据
        n, remoteAddr, err := pConn.ReadFrom(data)
        if err != nil {
			Logger.WarnLog(err)
            continue
        }

		b := make([]byte,n)
		copy(b,data)
		Logger.DebugLog("received ",b," from ",remoteAddr)
		packetConn := &UdpPacketConn{
			addr : remoteAddr,
			conn : pConn,
		}

		server.waitGroup.Add(1)
		go u.handleUdpPacket(h,packetConn,b)
    }

}
