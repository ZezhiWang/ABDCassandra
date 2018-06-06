package main 

import(
	"os"
	"fmt"
	"bufio"
	"strings"
)

func client(){
	reader := bufio.NewReader(os.Stdin)
    for {
	    fmt.Print("->")
        // handle command line input
        text,_ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)
        if strings.HasPrefix(text, "write"){
            input := strings.SplitN(text, " ", 3)
            key := input[1]
            // abd write
            write(key, []byte(input[2]))
        } else if strings.HasPrefix(text, "read") {
            key := strings.SplitN(text, " ", 2)[1]
            // output abd read
            fmt.Printf("\t%s\n",read(key))
        } else {
            // quit
            break
        }
    }
}
