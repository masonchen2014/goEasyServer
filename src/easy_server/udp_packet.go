package easy_server

import(
    "net"
    "sync"
)

type UdpPacketConn struct{
     addr net.Addr
     conn net.PacketConn
}

func (p * UdpPacketConn) SendData(bytes []byte){
     _,err := p.conn.WriteTo(bytes,p.addr)
     if err != nil {
     	Logger.WarnLog(err)
     }
}

type UdpPacketHandler struct{
     waitGroup * sync.WaitGroup
}

func (u * UdpPacketHandler)handleUdpPacket(handler func(UdpPacketOps,[]byte),conn * UdpPacketConn,bytes []byte){
    defer u.waitGroup.Done()
    handler(conn,bytes)
}

type UdpPacketOps interface{
     SendData([]byte)
}
