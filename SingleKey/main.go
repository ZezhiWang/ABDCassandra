package main

import(
	"os"
	"fmt"
	"flag"
	"bufio"
    "strings"
	"github.com/gocql/gocql"
)

//	Keyspace 	= demo
//	Table		= abd(key int, id text, ver int, val text)

var addrs = []string{"172.17.0.2", "172.17.0.3", "172.17.0.4"}

var (
	id 			int
	sessions 	map[string]*gocql.Session	
)

func main() {
	flag.IntVar(&id, "id", 0, "specify the node id")
	sessions = make(map[string]*gocql.Session)
	for _,addr := range addrs {
		sessions[addr] = getSession(addr)
	}
	defer closeAll()

	done := make(chan bool)
	go userInput(done)
	<-done
}

func userInput(done chan bool) {
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter Commands: ")
    for {
        text,_ := reader.ReadString('\n')
        if strings.HasPrefix(text, "write"){
        	text = strings.Replace(text, "\n", "", -1)
        	info := strings.SplitN(text, " ", 2)[1]
        	write(info)
        } else if strings.HasPrefix(text, "read") {
        	fmt.Printf("\t%s\n",read())
        } else {
        	break
        }
    }
    done <- true
}
