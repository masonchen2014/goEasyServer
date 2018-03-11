package easy_server

type receiver struct{
	dataFuncChan chan tcpDataFuncPacket
}

func (r * receiver) splitPacket(tcpConn *TcpConnection,h * TcpDataHandlers){
	defer tcpConn.Close()
	conn := tcpConn.conn
	bytes := make([]byte,512)
	bufferOffset := 0
	sendBufferOffset := 0
	var bytesToSend	[]byte
	lastPacketRemainBytes := 0
	bytesToHandleOffset := 0
	firstPacket := true
	remoteAddr :=conn.RemoteAddr()

	for{
		//here handle the last uncomplete packet
		if lastPacketRemainBytes > 0 {
			n,err := conn.Read(bytesToSend[sendBufferOffset:])
			if err !=nil {
			    Logger.ErrorLog(err)
				break
			}
			lastPacketRemainBytes -= n
			sendBufferOffset += n

			//here handle the complete packet
			if lastPacketRemainBytes ==0 {
				if firstPacket !=true{
					r.dataFuncChan <- tcpDataFuncPacket{tcpConn,bytesToSend,h}
				}else{
					firstPacket =false
					Logger.DebugLog("Here handles the first packet ",bytesToSend," for connection from ",remoteAddr)
					h.handleFirstPacket(tcpConn,bytesToSend)
				}
			}

			continue
		}

		//new packet
		n,err := conn.Read(bytes[bufferOffset:])
		if err !=nil {
			Logger.ErrorLog(err)
			break
		}
		bufferOffset +=n

		for bytesToHandleOffset < bufferOffset {
			//the buffer has some bytes now ,let's go handle them!
			if h.splitPacket !=nil {
				tcpPacketLen,err := h.splitPacket(bytes[bytesToHandleOffset:bufferOffset])
				if err == LessDataSplitError {
					Logger.WarnLog("Data is not enough to analysis its length for conneciton from ",remoteAddr)
					break
				}

				if err == OtherSplitError {
					Logger.WarnLog("Can't ananysis the length ,now handle all the data received from ",remoteAddr)
					//here handle all the bytes in buffer
					tcpPacketLen = bufferOffset
					bytesToHandleOffset = bufferOffset
					break
				}

				if tcpPacketLen<=0 {
					Logger.ErrorLog("Your split package func must have something wrong! go check it!!!")
					return;
				}

				//assign not declare
				bytesToSend = make([]byte,tcpPacketLen)
				copy(bytesToSend,bytes[bytesToHandleOffset:bufferOffset])

				if tcpPacketLen > bufferOffset-bytesToHandleOffset {
					//here means all the bytes in the buffer copied to the send buffer
					sendBufferOffset = bufferOffset-bytesToHandleOffset
					lastPacketRemainBytes = tcpPacketLen-sendBufferOffset
					bytesToHandleOffset = bufferOffset
					bufferOffset = 0
				}else{
					//here means there are remain bytes in buffer
					bytesToHandleOffset += tcpPacketLen
					//here send the complete packet
					// go HandleTcpPacket(server.tcp_packet_channel,)
					if firstPacket !=true{
						r.dataFuncChan <- tcpDataFuncPacket{tcpConn,bytesToSend,h}
					}else{
						firstPacket =false
						h.handleFirstPacket(tcpConn,bytesToSend)
						Logger.DebugLog("Here handles the first packet ",bytesToSend," for connection from ",remoteAddr)
					}
				}

			}else{
				bytesToSend = make([]byte,n)
				copy(bytesToSend,bytes[bytesToHandleOffset:bufferOffset])
				bytesToHandleOffset = 0
				bufferOffset = 0

				if firstPacket !=true{
					r.dataFuncChan <- tcpDataFuncPacket{tcpConn,bytesToSend,h}
				}else{
					Logger.DebugLog("Here handles the first packet ",bytesToSend," for connection from ",remoteAddr)
					firstPacket =false
					h.handleFirstPacket(tcpConn,bytesToSend)
				}
			}

		}

	}


}
