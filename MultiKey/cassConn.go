package main

import(
	"fmt"
	"log"
	"github.com/gocql/gocql"
)

func getSession(addr string) *gocql.Session {
	cluster := gocql.NewCluster(addr)
	cluster.Keyspace = "demo"
	cluster.Consistency = "One"
	session,_ := cluster.CreateSession()
	return session
}

func closeAll() {
	for _, sess := range sessions {
		sess.Close()
	}
	fmt.Println("all sessions closed")
}

func queryGet(key int, session *gocql.Session, done chan bool) {
	var tv TagVal
	arg := fmt.Sprintf("SELECT id, ver, val FROM tmp WHERE key=%d", key)
	if err := session.Query(arg).Scan(&tv.Id, &tv.Ver, &tv.Val); err != nil {
		log.Fatal(err)
	}
	done <-tv
}

func querySet(key int, tv TagVal, session *gocql.Session, done chan bool) {
	// update node tag
	arg := fmt.Sprintf("UPDATE abd SET id='%s', ver=%d, WHERE key=0", tv.Id, tv.Ver)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
	// insert value
	arg = fmt.Sprintf("INSERT INTO abd (key, id, ver, val) values (%d, %s, %d, %s)", key, tv.Id, tv.Ver, tv.Val)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
	done <-true
}