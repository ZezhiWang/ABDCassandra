package main

func write(val string){
	tv := get()
	tv.update(addrs[id], val)
	set(tv)
}

func read() string{
	tv := get()
	set(tv)
	return tv.Val
}

func get() TagVal{
	done := make(chan TagVal)
	for _,session := range sessions {
		go queryGet(session, done)
	}
	
	tv := TagVal{"", 0, ""}
	for i := 0; i < len(sessions)/2 + 1; i++ {
		tmp := <-done
		if tv.smaller(tmp) {
			tv = tmp
		}
	}
	return tv
}

func set(tv TagVal){
	done := make(chan bool)
	for _,session := range sessions {
		go querySet(tv, session, done)
	}
	
	for i := 0; i < len(sessions)/2 + 1; i++ {
		<-done
	}
}
