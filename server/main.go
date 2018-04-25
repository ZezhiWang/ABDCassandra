package main 

import(
	"github.com/gocql/gocql"
)

var (
	port	string
	cassIP  string
	state 	Tag
	session *gocql.Session
)

const GET=0
const SET=1

func main(){
	flag.StringVar(&port, "clientID", "5001", "input client ID")
	flag.StringVar(&cassIP, "cassIP", "172.0.0.2", "input cassIP")
	flag.Parse()
	server_task()
}
