package main 

import(
	"os"
	"fmt"
	"bufio"
	"strings"
)

func client(){
	reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter Commands: ")
    for {
	fmt.Print("->")
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
}
