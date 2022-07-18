package main

import (
	"net"
	"time"

	"github.com/jcbritobr/gonet/helper"
)

func main() {
	service := ":1200"
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	helper.CheckError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	helper.CheckError(err)
	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	helper.CheckError(err)
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
