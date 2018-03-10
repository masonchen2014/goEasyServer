package main
import(
	"easy_server"
	"fmt"
	"os"
	"time"
	// "runtime"
)

func JustPrint(op easy_server.TcpConnectionOps,bytes []byte){
     fmt.Println(bytes)
	 op.SendData(bytes)
}

func CloseTheConnection(op easy_server.TcpConnectionOps,bytes []byte){
     time.Sleep(10*time.Second)
     op.Close()

}

func UdpDataHandler(ops easy_server.UdpPacketOps,bytes []byte){
     fmt.Println(bytes)
     ops.SendData(bytes)
}								//

func main(){
	 // runtime.GOMAXPROCS(1)
	file,_ := os.Create("dxp.log")
	easy_server.SetEasyLogger(os.Stdout,file,file,file)
	server := easy_server.NewServer()
	// server.CreateWorkers()
	handlers := easy_server.NewTcpDataHandlers(nil,JustPrint,JustPrint,2) //
	// time.Sleep(10*time.Second)
	server.AddTcpListener(":4003",handlers)

	// server.AddUdpListener(":4003",UdpDataHandler)
	go server.PrintServerInfo()
	// time.Sleep(100*time.Second)

	server.Stop()
}
