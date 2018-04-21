package main 

import (
	zmq "github.com/"
)

func createDealerSocket() *zmq.Socket {
	dealer,_ := zmq.NewSocket(zmq.DEALER)
	var addr string
	for _,server := range servers {
		addr = "tcp://" + server + ":8888"
		dealer.Connect(addr)
		fmt.Println(addr)
	}
	return dealer
}

func sendToServer(msg Message, dealer *zmq.Socket) {
	msgToSend := getGobFromMsg(msg)
	dealer.SendBytes(msgToSend.Bytes(), NON_BLOCKING)
}

func recvData(dealer *zmq.Socket) TagVal {
	msgBytes,_ := dealer.RecvBytes(0)
	msg := getMsgFromGob(msgBytes)
	if msg.OpType != GET {
		return recvData()
	}
	return msg.Tv
}

func recvData(dealer *zmq.Socket) TagVal {
	msgBytes,_ := dealer.RecvBytes(0)
	msg := getMsgFromGob(msgBytes)
	if msg.OpType != SET {
		recvData(dealer)
	}
}