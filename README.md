# goEasyServer
A high performance tcp/udp server written by golang and it is very simple to use,the belowing are the steps :

first step: import the package "easy_server"

      import(
         "easy_server"
      )
      
second step: create the EasyServer instance with the parameter worker number

      server := easy_server.NewServer(8)
      
third step: create the workers using the number of workers

      server.CreateWorkers()
      
forth step: create data handlers,one for spliting the data,one for handling the first splited packet ,and the last is used to hangle the non-first packets

      handlers := easy_server.NewTcpDataHandlers(splitFunc,handleFirstPacket,handleNonFirstPacket)
      
fifth step: add tcp listenner on specified port

      server.AddTcpListener(":4003",handlers)
       
last step: wait for all the jobs done

      server.Stop()
      

log handling: by default all the logs will be output to the os.Stdout, you can define your own log files as below

     file,_ := os.Create("server.log")
     easy_server.SetEasyLogger(os.Stdout,file,file,file)
