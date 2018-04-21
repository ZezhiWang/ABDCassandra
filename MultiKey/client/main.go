package main

import "flag"

//	Keyspace 	= demo
//	Table		= abd(key int, id text, ver int, val text)

var (
	ID string
	servers = []string{"172.17.0.2", "172.17.0.3", "172.17.0.4"}
	addrs 

)

func main() {
	flag.StringVar(&ID, "clientID", "172.17.0.1", "input client ID")
	flag.Parse()		

	client()
}
