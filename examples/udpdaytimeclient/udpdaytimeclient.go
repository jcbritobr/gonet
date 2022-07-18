package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jcbritobr/gonet/helper"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage:", os.Args[0], "host:port")
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	helper.CheckError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	helper.CheckError(err)
	_, err = conn.Write([]byte("anything"))
	helper.CheckError(err)
	var buf [512]byte
	n, err := conn.Read(buf[0:])
	helper.CheckError(err)
	fmt.Println(string(buf[0:n]))
}
