package main
import (
	"net"
	"os"
    "fmt"
    "bufio"
    "strings"
)
func main() {
	addr := net.ParseIP("127.0.0.1")
    if addr == nil {
        fmt.Println("Invalid address")
    } else {
        fmt.Println("The address is ", addr.String())
	}
	
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "")
	checkError(err)
    fmt.Println("The address is ", tcpAddr.String())
    
    ln, err := net.Listen("tcp", ":9090")
    if(err != nil){
        fmt.Println(err)
    } else {

    for {
        conn, err := ln.Accept()
        if err != nil {
            // handle error
        }
        go handleConnection(conn)
    }
}
   
}

func handleConnection( conn net.Conn) {
    message, err := bufio.NewReader(conn).ReadString('\n')
    if(err != nil){
        fmt.Println(err)
    } else {
        fmt.Print("Message Received:", string(message))
        newmessage := strings.ToUpper(message)
        // send new string back to client
        conn.Write([]byte(newmessage + "\n"))
    }

}
func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}