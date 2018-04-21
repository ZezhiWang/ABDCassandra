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

// func closeAll() {
// 	for _, server := range servers {
// 		server.session.Close()
// 	}
// 	fmt.Println("all servers closed")
// }

func queryGet(key string) string {
	var res string
	arg := fmt.Sprintf("SELECT val FROM abd WHERE key='%s'", key)
	if err := session.Query(arg).Scan(&res); err != nil {
		log.Fatal(err)
	}
	return res
}

func querySet(tv TagVal) {
	arg := fmt.Sprintf("UPDATE abd SET Id='%s', Val='%s', Ts=%d WHERE Key='%s'", tv.Id, tv.Ts, tv.Val, tv.Key)
	if err := session.Query(arg).Exec(); err != nil {
		log.Fatal(err)
	}
}
