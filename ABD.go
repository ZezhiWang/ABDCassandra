package main

import(
	"fmt"
)

func write(val string){
	tv := get()
	tv = tv.update(addr, val)
	set(tv)
}

func read(){
	tv := get()
	set(tv)
	return tv.Val
}

func get() Tagval{
	done := make(chan TagVal)
	for _,session := range sessions {
		go queryGet(session, done)
	}
	
	tv := TagVal{"", 0, ""}
	for i := 0; i < len(sessions)/2 + 1; i++ {
		tmp := <-done
		if tag.smaller(tmp) {
			tv = tmp
		}
	}
	return tv
}

func set(key int, tv Tagval){
	done := make(chan bool)
	for _,session := range sessions {
		go queryPut(tv, session, done)
	}
	
	for i := 0; i < len(sessions)/2 + 1; i++ {
		<-done
	}
}