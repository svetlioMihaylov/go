package main
import (
	"net"
	"time"
	"fmt"
	 "bufio"
    
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
    if(err != nil){
        fmt.Println(err)
    } else {

		fmt.Fprintf(conn, "test" + "\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: "+message)

	}
	time.Sleep(2 * time.Second)
}