package main

import "fmt"

func write(key int, val string){
	tv := get(0)
	tv.update(id, val)
//	fmt.Println(tv)
	set(key, tv)
}

func read(key int) string{
	tv := get(key)
//	fmt.Println(tv)
	set(key,tv)
	return tv.Val
}

func get(key int) TagVal{
	done := make(chan TagVal)
	for _,s := range servers {
		go s.getFromServer(key, done)
	}
	
	tv := TagVal{"", 0, ""}
	for i := 0; i < len(servers)/2 + 1; i++ {
		tmp := <-done
		fmt.Println(tmp)
		if tv.smaller(tmp) {
			tv = tmp
		}
	}
	return tv
}

func set(key int, tv TagVal){
	done := make(chan bool)
	for _,s := range servers {
		go s.setToServer(key, tv, done)
	}
	
	for i := 0; i < len(servers)/2 + 1; i++ {
		<-done
	}
}
