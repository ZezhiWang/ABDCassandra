package main

import "fmt"

func write(key int, val string){
	getChan := make(chan TagVal)
	go get(0, getChan)
	tv := <- getChan
	tv.update(id, val)
	setChan := make(chan bool)
	go set(key, tv, setChan)
	<- setChan
}

func read(key int) string{
	getChan := make(chan TagVal)
	go get(key, getChan)
	tv := <- getChan
	tv.update(id, val)
	setChan := make(chan bool)
	go set(key, tv, setChan)
	<- setChan
	return tv.Val
}

func get(key int, getChan chan TagVal) {
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
	
	getChan <- tv

	for i := len(servers)/2 + 1; i < len(servers); i++ {
		<-done
	}
}

func set(key int, tv TagVal, setChan chan bool){
	done := make(chan bool)
	for _,s := range servers {
		go s.setToServer(key, tv, done)
	}
	
	for i := 0; i < len(servers)/2 + 1; i++ {
		<-done
	}

	setChan <- true

	for i := len(servers)/2 + 1; i < len(servers); i++ {
		<-done
	}
}
