package main

import(
	"github.com/gocql/gocql"
)

//	Keyspace 	= demo
//	Table		= abd(key int, id text, ver int, val text)

var addrs = []string{"172.17.0.2", "172.17.0.3", "172.17.0.4"}

var (
	servers 	map[int]Server
)

func main() {
	servers = make(map[int]Server)
	for id,addr := range addrs {
		tv := TagVal{Id: addr, Ver: 0, Val: ""}
		servers[id] = Server{tag: tv, session: getSession(addr)}
	}
	defer closeAll()

	done := make(chan bool)
	client(done)
}