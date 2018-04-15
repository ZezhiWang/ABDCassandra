package main

func write(key int, val string){
	tv := get(0)
	tv.update(addrs[id], val)
	set(key, tv)
}

func read(key int) string{
	tv := get(key)
	set(key,tv)
	return tv.Val
}

func get(key int) TagVal{
	done := make(chan TagVal)
	for _,session := range sessions {
		go queryGet(key, session, done)
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

func set(key int, tv TagVal){
	done := make(chan bool)
	for _,session := range sessions {
		go querySet(key, tv, session, done)
	}
	
	for i := 0; i < len(sessions)/2 + 1; i++ {
		<-done
	}
}
