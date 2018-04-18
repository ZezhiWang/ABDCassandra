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
        text,_ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)
        if strings.HasPrefix(text, "write"){
            input := strings.SplitN(text, " ", 3)
            key,_ := strconv.Atoi(input[1])
            write(key, input[2])
        } else if strings.HasPrefix(text, "read") {
            key,_ := strconv.Atoi(strings.SplitN(text, " ", 2)[1])
            fmt.Printf("\t%s\n",read(key))
        } else {
            break
        }
    }
}
