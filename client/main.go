package main 

import "flag"

//	Keyspace 	= demo
//	Table		= abd(key text, id text, ver int, val text)

var (
	ID string
	// IP addresses of servers
	servers = []string{"127.0.0.1:5001", "127.0.0.1:5002", "127.0.0.1:5003"}	
)

// used to mark the phase
const GET=0
const SET=1

func main() {
	// init client id
	flag.StringVar(&ID, "clientID", "172.17.0.1", "input client ID")
	flag.Parse()		

	client()
}
