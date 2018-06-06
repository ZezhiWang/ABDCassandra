package main 

import (
	"fmt"
	"log"
	zmq "github.com/pebbe/zmq4"
)

func server_task() {
	// Set the ZMQ sockets
	frontend,_ := zmq.NewSocket(zmq.ROUTER)
	defer frontend.Close()
	frontend.Bind("tcp://*:"+port)

	//  Backend socket talks to workers over inproc
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer backend.Close()
	backend.Bind("inproc://backend")

	go server_worker()

	//  Connect backend to frontend via a proxy
	err := zmq.Proxy(frontend, backend, nil)
	log.Fatal("Proxy interrupted:", err)
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
		if err != nil {
			fmt.Println(err)
		}
		// decode message
		message := getMsgFromGob(msg[1])
		msg_reply[0] = msg[0]

		// create response message
		tmpMsg := createRep(message)
		// encode message
		tmpGob := getGobFromMsg(tmpMsg)
		msg_reply[1] = tmpGob.Bytes()

		worker.SendMessage(msg_reply)
	}
}

// create response message
func createRep(input Message) Message {
	var output Message
	switch input.OpType{
	// if set phase
	case SET:
		querySet(input.Tv)
		output.OpType = SET
		output.Tv = TagVal{Tag: state, Key: "", Val: make([]byte, 4)}
	// if get phase
	case GET:
		output.OpType = GET
		output.Tv = queryGet(input.Tv.Key)
	}
	return output
}
