package main 

import "flag"

//	Keyspace 	= demo
//	Table		= abd(key text, id text, ver int, val text)

var (
	ID string
	// IP addresses of servers
	servers = []string{"128.52.162.127:5001", "128.52.162.122:500`", "128.52.162.123:5001"}	
)

// used to mark the phase
const GET=0
const SET=1

func main() {
	// init client id
	flag.StringVar(&ID, "clientID", "128.52.162.120", "input client ID")
	flag.Parse()		

	client()
}
