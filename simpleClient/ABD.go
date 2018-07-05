package main 

// abd write
func write(key string, val []byte){
	tv := get(key)
	tv.update(ID, val)
	set(tv)
}

// abd read
func read(key string) []byte{
	tv := get(key)
	set(tv)
	return tv.Val
}

// get phase
func get(key string) TagVal {
	done := make(chan TagVal)
	for _,session := range servers {
		go queryGet(key, session, done)
	}

	// init tagval
	tv := TagVal{Tag: Tag{Id: "0", Ts: 0}, Key: key, Val: make([]byte,4)}
	// find tagval with biggest tag in quorum 
	for i := 0; i < len(servers)/2 + 1; i++ {
		tmp := <- done
		if tv.Tag.smaller(tmp.Tag) {
			tv = tmp
		}
	}
	return tv	
}

// set phase
func set(tv TagVal){
	done := make(chan bool)
	for _,session := range servers {
		go querySet(tv, session, done)
	}

	// recv ack from quorum
	for i := 0; i < len(servers)/2 + 1; i++ {
		<-done
	}
}
