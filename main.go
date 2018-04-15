package main

import(
	"github.com/gocql/gocql"
)

"""
	Keyspace 	= demo
	Table		= abd(key int, id text, ver int, val text)
	IP			= 172.17.0.1
"""

var (
	id 			int
	addrs 		[]string
	sessions 	map[string]*gocql.Session	
)

func main() {
	init()
	defer closeAll()
}

func init() {
	addrs = []string{"172.17.0.1"}
	for _, addr := range addrs {
		sessions[addr] = getSession(addr)
	}
}