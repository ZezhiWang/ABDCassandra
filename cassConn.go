package main

import(
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
	for _, sess := range sessions {
		sess.Close()
	}
}

func queryGet(session *gocql.Session, done chan bool) {
	var tv TagVal
	arg := fmt.Sprintf("SELECT id, ver, val FROM tmp WHERE key=0")
	if err := session.Query(arg).Scan(&tv.Id, &tv.Ver, &tv.val); err != nil {
		log.Fatal(err)
	}
	done <- tv
}

func querySet(tv TagVal, session *gocql.Session, done chan bool) {
	arg := fmt.Sprintf("UPDATE abd SET id=%s, ver=%d, val=%s WHERE key=0", tv.Id, tv.Ver, tv.Val)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
	done <-true
}