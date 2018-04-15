package main

func write(key int, val string){
	tv := get(0)
	tv.update(addrs[id], val)
	set(key, tv)
}

func read(key int){
	tv := get(key)
	set(key,tv)
	return tv.Val
}

func get(key int) Tagval{
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

func set(key int, tv Tagval){
	done := make(chan bool)
	for _,session := range sessions {
		go queryPut(key, tv, session, done)
	}
	
	for i := 0; i < len(sessions)/2 + 1; i++ {
		<-done
	}
}