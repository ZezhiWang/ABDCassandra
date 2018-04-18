package main

import "flag"

//	Keyspace 	= demo
//	Table		= abd(key int, id text, ver int, val text)

var (
	id	string
	servers 	map[int]Server
	addrs = []string{"172.17.0.2", "172.17.0.3", "172.17.0.4"}
)

func main() {
	flag.StringVar(&id, "clientID", "172.17.0.1", "input client ID")
	flag.Parse()		
	servers = make(map[int]Server)
	for id,addr := range addrs {
		tv := TagVal{Id: "", Ver: 0, Val: ""}
		servers[id] = Server{id: id, tag: tv, session: getSession(addr)}
	}
	defer closeAll()

	client()
}
