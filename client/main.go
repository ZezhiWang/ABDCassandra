package main 

import (
	"fmt"
	"flag"
	"sync"
	"time"
)

//	Keyspace 	= demo
//	Table		= abd(key text, id text, ver int, val text)

var (
	ID string
	mutex = &sync.Mutex{}
	// IP addresses of servers
//	servers = []string{"128.52.162.120:5001","128.52.162.128:5001","128.52.162.129:5001","128.52.162.127:5001", "128.52.162.122:5001", "128.52.162.123:5001","128.52.162.124:5001","128.52.162.125:5001","128.52.162.131:5001"}	
	servers = []string{"128.52.162.127:5001", "128.52.162.122:5001", "128.52.162.123:5001"}	
	data_size int
)

// used to mark the phase
const GET=0
const SET=1

func main() {
	// init client id
	flag.StringVar(&ID, "clientID", "128.52.162.120", "input client ID")
	flag.IntVar(&data_size, "dataSize", 1024, "input data size")
	flag.Parse()		

//	client()
	test()
}

func test(){
	num := 1000
	wTime := make(chan time.Duration)
	rTime := make(chan time.Duration)
	var WTotal, RTotal int = 0,0

//	for i := 0; i < num; i++{
//		write(string(i), make([]byte,1024))
//	}

	s := time.Now()	

	for i := 0; i < num; i++ {
		go testW(string(i),wTime)
		go testR(string(i),rTime)
	}

	for i := 0; i < num; i++ {
		WTotal += int(<-wTime/time.Millisecond)
		RTotal += int(<-rTime/time.Millisecond)
	}

	e := time.Now()
	t := e.Sub(s)

	fmt.Printf("Avg write time: %f ms\n", float64(WTotal)/float64(num))
	fmt.Printf("Avg read time: %f ms\n", float64(RTotal)/float64(num))
	fmt.Printf("Total time: %f ms\n",int(t/time.Millisecond))
}

func testW(key string, wTime chan time.Duration){
	mutex.Lock()
	start := time.Now()
	write(key,make([]byte,data_size))
	end := time.Now()
	mutex.Unlock()
	wTime <- end.Sub(start)
}

func testR(key string, rTime chan time.Duration){
	mutex.Lock()
	start := time.Now()
	read(key)
	end := time.Now()
	mutex.Unlock()
	rTime <- end.Sub(start)
}

	
