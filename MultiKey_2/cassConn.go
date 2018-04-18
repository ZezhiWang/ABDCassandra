package main

import(
	"fmt"
	"log"
	"github.com/gocql/gocql"
)

func getSession(addr string) *gocql.Session {
	cluster := gocql.NewCluster(addr)
	cluster.Keyspace = "demo"
	cluster.Consistency = gocql.One
	session,_ := cluster.CreateSession()
	return session
}

func closeAll() {
	for _, server := range servers {
		server.session.Close()
	}
	fmt.Println("all servers closed")
}

func queryGet(key int, session *gocql.Session) TagVal {
	var tmp TagVal
	var idx int
	arg := fmt.Sprintf("SELECT key, id, ver, val FROM abd")
	iter := session.Query(arg).Iter()
	tv := TagVal{"",0,""}
	for iter.Scan(&idx, &tmp.Id, &tmp.Ver, &tmp.Val) {
		if tv.smaller(tmp){
			tv.Id = tmp.Id
			tv.Ver = tmp.Ver
		}
		if idx == key {
			tv.Val = tmp.Val
		}
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return tv
}

func querySet(key int, tv TagVal, session *gocql.Session) {
	arg := fmt.Sprintf("INSERT INTO abd (key, id, ver, val) values (%d, '%s', %d, '%s')", key, tv.Id, tv.Ver, tv.Val)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
}
