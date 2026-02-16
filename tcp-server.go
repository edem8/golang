// Building a tcp socker server
package main

import (
	"net"
	"fmt"
	"log"
	"bufio"
	"strings"
)


func main(){

	// net.Listen starts a tcp server on a given network(TCP) and port(8090 on all interfaces)
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	} 


	// close the listner to free the port when the application exits
	defer listener.Close()


	// loop infintely to accept new client connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connections:", err)
		}

		go handleConnection(conn)
	}

}




func handleConnection(conn net.Conn){
	// closing the connection release the resource when interaction with client is complete
	defer conn.Close()

	// use bufio.NewReader to read data from the connection-one line
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Error reading from connection:", err)
		return
	}

	// create and send back response to acknowledge the reciever messsage
	ackMsg := strings.ToUpper(strings.TrimSpace(message))
	response := fmt.Sprintf("ACK: %s\n", ackMsg)
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error:", err)
		
	}

}