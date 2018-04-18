package main

func write(val string){
	getChan := make(chan TagVal)
	go get(getChan)
	tv := <- getChan
	tv.update(id, val)
	setChan := make(chan bool)
	go set(tv, setChan)
	<- setChan
}

func read() string{
	getChan := make(chan TagVal)
	go get(getChan)
	tv := <- getChan
	setChan := make(chan bool)
	go set(tv, setChan)
	<- setChan
	return tv.Val
}

func get(getChan chan TagVal) {
	done := make(chan TagVal)
	for _,s := range servers {
		go s.getFromServer(done)
	}
	
	tv := TagVal{"", 0, ""}
	for i := 0; i < len(servers)/2 + 1; i++ {
		tmp := <-done
		if tv.smaller(tmp) {
			tv = tmp
		}
	}
	getChan <- tv

	for i := len(servers)/2 + 1; i < len(servers); i++ {
		<-done
	}
}

func set(tv TagVal, setChan chan bool){
	done := make(chan bool)
	for _,s := range servers {
		go s.setToServer(tv, done)
	}
	
	for i := 0; i < len(servers)/2 + 1; i++ {
		<-done
	}

	setChan <- true

	for i := len(servers)/2 + 1; i < len(servers); i++ {
		<-done
	}
}
