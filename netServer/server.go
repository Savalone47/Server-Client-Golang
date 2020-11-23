package main

// run the Server and client, wait for the response
import (
	"fmt"
	"net"
	"os"
)
// Declare and assign constante
const (
	SvrHost = "localhost"
	SvrPort = "9982"
	SvrType = "tcp"
)
func main() {

	// Display the message for connection
	fmt.Println("Server is running ...")

	// listen to the server
	svr, err := net.Listen(SvrType, SvrHost+":"+SvrPort)

	if err != nil {
		fmt.Println("error listing ", err.Error())
		os.Exit(1)
	}

	defer svr.Close() // close the connection
	
	fmt.Println("Listing on " + SvrHost + ":" + SvrPort) // server success
	fmt.Println("waiting client....")                    // waiting for response of client

	// Accepting and connecting the client
	for  {
		conn, err := svr.Accept()
		if err!=nil {
			fmt.Println("error accepting ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Client connected :) ")
		go handleClient(conn) // call the function handleClient
	}
}

func handleClient(conn net.Conn){
	buff := make([] byte, 1024)
	msgLen, err := conn.Read(buff)
	if err!=nil {
		fmt.Println("Error reading ", err.Error())
	}
	fmt.Println("Received: ", string(buff[:msgLen]))
	_ = conn.Close()
}