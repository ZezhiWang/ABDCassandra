package main 

import (
//	zmq "github.com/pebbe/zmq3"
)

func write(key string, val string){
	tv := get(key)
	tv.update(ID, val)
	set(tv)
}

func read(key string) string{
	tv := get(key)
	fmt.Println("get phase completed")
	set(tv)
	return tv.Val
}

func get(key string) TagVal {
	dealer := createDealerSocket()
	defer dealer.Close()

	tv := TagVal{Tag: Tag{Id: "", Ts: 0}, Key: key, Val: ""}
	msg := Message{OpType: GET, Tv: tv}
	sendToServer(msg, dealer)
	fmt.Println("send completed")
	for i := 0; i < len(servers)/2 + 1; i++ {
		tmp := recvData(dealer)
		fmt.Println("recv one tag")
		if tv.Tag.smaller(tmp.Tag) {
			tv = tmp
		}
	}
	return tv	
}

func set(tv TagVal){
	dealer := createDealerSocket()
	defer dealer.Close()

	msg := Message{OpType: SET, Tv: tv}
	go sendToServer(msg, dealer)
	
	for i := 0; i < len(servers)/2 + 1; i++ {
		recvAck(dealer)
	}
}
