package main

import(
	"fmt"
	"github.com/gocql/gocql"
)

//	Keyspace 	= demo
//	Table		= abd(key int, id text, ver int, val text)
//	IP			= 172.17.0.1

var (
	id 			int
	addrs 		[]string
	sessions 	map[string]*gocql.Session	
)

func main() {
	addrs = []string{"172.17.0.2"}
	id = 0
	sessions = make(map[string]*gocql.Session)
	for _,addr := range addrs {
		sessions[addr] = getSession(addr)
	}
	defer closeAll()
	write("test")
	fmt.Println(read())
}
