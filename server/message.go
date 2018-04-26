package main

import (
	"fmt"
	"bytes"
	"encoding/gob"	
)

type Message struct {
	OpType 	int
	Tv 		TagVal
}

func getGobFromMsg(msg Message) bytes.Buffer {
	var res bytes.Buffer

	enc := gob.NewEncoder(&res)
	err := enc.Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func getMsgFromGob(msgBytes []byte) Message {
	var buff bytes.Buffer
	var msg Message

	buff.Write(msgBytes)
	dec := gob.NewDecoder(&buff)
	dec.Decode(&msg)

	return msg
}
