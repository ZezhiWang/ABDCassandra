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

func queryGet(key string) string {
	var res string
	session := getSession(cassIP)
	defer session.Close()
	arg := fmt.Sprintf("SELECT val FROM abd WHERE key='%s'", key)
	if err := session.Query(arg).Scan(&res); err != nil {
		log.Fatal(err)
	}
	return res
}

func querySet(tv TagVal) {
	session := getSession(cassIP)
	defer session.Close()
	arg := fmt.Sprintf("UPDATE abd SET Id='%s', Val='%s', Ts=%d WHERE Key='%s'", tv.Tag.Id, tv.Val, tv.Tag.Ts, tv.Key)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
}
