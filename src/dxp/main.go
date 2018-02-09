package main
import(
	"easy_server"
	"fmt"
)

func JustPrint(op easy_server.TcpConnectionOps,bytes []byte){
     fmt.Println(bytes)
}

func main(){
	// runtime.GOMAXPROCS(1)
	server := easy_server.NewServer(8)
	server.CreateWorkers()
	server.PrintServerInfo()
	handlers := easy_server.NewTcpDataHandlers(nil,JustPrint,JustPrint)
	server.AddTcpListener(":4003",handlers)
	server.Stop()
}
