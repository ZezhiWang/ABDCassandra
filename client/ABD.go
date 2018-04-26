package main 

// abd write
func write(key string, val string){
	tv := get(key)
	tv.update(ID, val)
	set(tv)
}

// abd read
func read(key string) string{
	tv := get(key)
	set(tv)
	return tv.Val
}

// get phase
func get(key string) TagVal {
	dealer := createDealerSocket()
	defer dealer.Close()

	msg := Message{OpType: GET, Tv: tv}
	sendToServer(msg, dealer)
	
	// init tagval
	tv := TagVal{Tag: Tag{Id: "0", Ts: 0}, Key: key, Val: ""}
	// find tagval with biggest tag in quorum 
	for i := 0; i < len(servers)/2 + 1; i++ {
		tmp := recvData(dealer)
		if tv.Tag.smaller(tmp.Tag) {
			tv = tmp
		}
	}
	return tv	
}

// set phase
func set(tv TagVal){
	dealer := createDealerSocket()
	defer dealer.Close()

	msg := Message{OpType: SET, Tv: tv}
	sendToServer(msg, dealer)
	// recv ack from quorum
	for i := 0; i < len(servers)/2 + 1; i++ {
		recvAck(dealer)
	}
}
