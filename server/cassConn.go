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

// query cassandra to get val with key
func queryGet(key string) TagVal {
	var tv TagVal
	tv.Key = key
	arg := fmt.Sprintf("SELECT val, id, ver FROM abd WHERE key='%s'", key)
	if err := session.Query(arg).Scan(&tv.Val, &tv.Tag.Id, &tv.Tag.Ts); err != nil {
		//log.Fatal(err)
		tv.Val = nil
		tv.Tag.Id = ""
		tv.Tag.Ts = 0
	}
	return tv
}

// update tagval to cassandra
func querySet(tv TagVal) {
	arg := fmt.Sprintf("UPDATE abd SET id='%s',val=?,ver=%d WHERE key=? IF ver>%d",tv.Tag.Id,tv.Tag.Ts,tv.Tag.Ts)
	fmt.Println(arg)
	if err := session.Query(arg, tv.Val,tv.Key).Exec(); err != nil {
		log.Fatal(err)
	}
}
