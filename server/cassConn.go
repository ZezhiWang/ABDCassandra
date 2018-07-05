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
	arg := fmt.Sprintf("INSERT INTO abd (key,id,val,ver) values (?,?,?,?) WHERE ver < %d",tv.Tag.Ts)
	if err := session.Query(arg, tv.Key,tv.Tag.Id,tv.Val,tv.Tag.Ts).Exec(); err != nil {
		log.Fatal(err)
	}
}
