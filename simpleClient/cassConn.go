package main

import(
	"fmt"
	"log"
	"github.com/gocql/gocql"
)

// get cassandra session
func getSession(addr string) *gocql.Session {
	cluster := gocql.NewCluster(addr)
	cluster.Keyspace = "demo"
	cluster.Consistency = gocql.One
	session,err  := cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
	}
	return session
}

func closeAll() {
	for _, sess := range servers {
		sess.Close()
	}
	fmt.Println("all sessions closed")
}

// query cassandra to get val with key
func queryGet(key string, session *gocql.Session, done chan TagVal) {
	var tv TagVal
	tv.Key = key
	arg := fmt.Sprintf("SELECT val, id, ver FROM abd WHERE key='%s'", key)
	if err := session.Query(arg).Scan(&tv.Val, &tv.Tag.Id, &tv.Tag.Ts); err != nil {
		tv.Val = nil
		tv.Tag.Id = ""
		tv.Tag.Ts = 0
	}
	done <- tv
}

// update tagval to cassandra
func querySet(tv TagVal, session *gocql.Session, done chan bool) {
	arg := fmt.Sprintf("INSERT INTO abd (key,id,val,ver) values (?,?,?,?) WHERE ver < ?")
	if err := session.Query(arg, tv.Key,tv.Tag.Id,tv.Val,tv.Tag.Ts, tv.Tag.Ts).Exec(); err != nil {
		log.Fatal(err)
	}
}
