package main 

import(
	"flag"
)

var (
	port	string
	cassIP  string
	state 	Tag
)

const GET=0
const SET=1

func main(){
	flag.StringVar(&port, "clientID", "5001", "input client ID")
	flag.StringVar(&cassIP, "cassIP", "172.0.0.2", "input cassIP")
	flag.Parse()
	// Set the Default State Variables
	state = Tag{Id: "", Ts: 0}
	server_task()
}
