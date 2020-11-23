package main

// run client and wait for the response
import (
	"fmt"
	"net"
)

const (
	SvrHost = "localhost"
	SvrPort = "9982"
	SvrType = "tcp"
)
func main() {
	fmt.Println("Client is running...")
	conn, err := net.Dial(SvrType, SvrHost+":"+SvrPort)
	if err!=nil {
		panic(err)
	}
	fmt.Println("send data")
	_,err = conn.Write([]byte("message from client"))
	defer conn.Close()
}
