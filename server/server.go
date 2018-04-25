package main 

import (
	"fmt"
	"log"
//	"github.com/gocql/gocql"
	zmq "github.com/pebbe/zmq4"
)

func server_task() {
	// Set the Default State Variables
	state = Tag{Id: "", Ts: 0}
	session = getSession(cassIP)
	defer session.Close()
	// Set the ZMQ sockets
	frontend,err := zmq.NewSocket(zmq.ROUTER)
	if err != nil {
		fmt.Println(err)
	}
	defer frontend.Close()
	frontend.Bind("tcp://*:"+port)

	//  Backend socket talks to workers over inproc
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	backend.Bind("inproc://backend")

	go server_worker()

	//  Connect backend to frontend via a proxy
	err2 := zmq.Proxy(frontend, backend, nil)
	log.Fatal("Proxy interrupted:", err2)
}

func server_worker() {
	worker, _ := zmq.NewSocket(zmq.DEALER)
	defer worker.Close()
	worker.Connect("inproc://backend")
	msg_reply := make([][]byte, 2)

	for i := 0; i < len(msg_reply); i++ {
		msg_reply[i] = make([]byte, 0) // the frist frame  specifies the identity of the sender, the second specifies the content
	}

	for {
		msg,err := worker.RecvMessageBytes(0)
		fmt.Println(msg)
		if err != nil {
			fmt.Println(err)
		}

		message := getMsgFromGob(msg[1])
		msg_reply[0] = msg[0]
		fmt.Println(msg[0])

		tmpMsg := createRep(message)
		tmpGob := getGobFromMsg(tmpMsg)
		msg_reply[1] = tmpGob.Bytes()

		worker.SendMessage(msg_reply)
	}
}

func createRep(input Message) Message {
	var output Message
	switch input.OpType{
	case SET:
		output.OpType = SET
		if state.smaller(input.Tv.Tag) {
			state = input.Tv.Tag
			querySet(input.Tv)
		}
		output.Tv = TagVal{Tag: state, Key: "", Val: ""}
	case GET:
		output.OpType = GET
		tv := input.Tv
		tv.Tag = state
		tv.Val = queryGet(input.Tv.Key)
		output.Tv = tv
	}
	return output
}
