package main 

import (
	"fmt"
	zmq "github.com/pebbe/zmq4"
)

func createDealerSocket() *zmq.Socket {
	dealer,_ := zmq.NewSocket(zmq.DEALER)
	var addr string
	for _,server := range servers {
		addr = "tcp://" + server
		fmt.Println(addr)
		dealer.Connect(addr)
	}
	return dealer
}

func sendToServer(msg Message, dealer *zmq.Socket) {
	msgToSend := getGobFromMsg(msg)
	dealer.SendBytes(msgToSend.Bytes(), 0)
}

func recvData(dealer *zmq.Socket) TagVal {
	msgBytes,_ := dealer.RecvBytes(0)
	msg := getMsgFromGob(msgBytes)
	if msg.OpType != GET {
		return recvData(dealer)
	}
	return msg.Tv
}

func recvAck(dealer *zmq.Socket) {
	msgBytes,_ := dealer.RecvBytes(0)
	msg := getMsgFromGob(msgBytes)
	if msg.OpType != SET {
		recvAck(dealer)
	}
}
