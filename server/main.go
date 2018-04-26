package main 

import(
	"flag"
	"github.com/gocql/gocql"
)

var (
	port	string
	cassIP  string
	state 	Tag
	session *gocql.Session
)

// used to mark the phase
const GET=0
const SET=1

func main(){
	// init port and cassandra storage IP
	flag.StringVar(&port, "clientID", "5001", "input client ID")
	flag.StringVar(&cassIP, "cassIP", "172.0.0.2", "input cassIP")
	flag.Parse()

	// Set the Default State Variables
	state = Tag{Id: "", Ts: 0}
	// create cassandra session
	session = getSession(cassIP)
	defer session.Close()

	server_task()
}
