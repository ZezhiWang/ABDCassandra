package func main

import (
	"fmt"
	"encoding/gob"	
)

type Message struct {
	OpType 	int
	Sender 	string
	Tv 		TagVal
}

func createMsg(OpType int, tv TagVal) Message {
	var msg Message
	msg.Tv = 
}

func getGobFromMsg(msg Message) bytes.Buffer {
	var res bytes.Buffer

	enc := gob.NewEncoder(&res)
	if err := enc.Encode(msg); err != nil {
		fmt.Println(err)
	}
	return res
}

func getMsgFromGob(msgBytes []byte) Message {
	var buff bytes.Buffer
	var msg Message

	buff.Write(msgBytes)
	dec := gob.NewDecoder(&buff)
	if err := dec.Decode(&msg); err != nil {
		fmt.Println(err)
	}
	return msg
}