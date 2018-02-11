package easy_server

import(
	"net"
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

type udpDataFuncPacket struct{
     conn * UdpPacketConn
     bytes []byte
     handler func(UdpPacketOps,[]byte)
}

type UdpPacketOps interface{
     SendData([]byte)
}