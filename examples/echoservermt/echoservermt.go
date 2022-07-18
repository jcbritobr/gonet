package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	service := ":1201"
	//build the addresss
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	// create listener using the address
	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		// accept incomming connections
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
		// conn.Close() move the close func inside handleClient
	}
}

func handleClient(conn net.Conn) {
	// close connection on func exit
	defer conn.Close()

	var buf [512]byte
	for {
		// read data from client
		n, err := conn.Read(buf[:])
		checkError(err)
		fmt.Println(string(buf[:]))
		// write data to client
		_, err = conn.Write(buf[:n])
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatalln("Fatal error", err.Error())
	}
}
