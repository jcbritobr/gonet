package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Default().Fatalf("Usage: %s dotted-ip-addr ones bits\n", os.Args[0])
	}
	dotAddr := os.Args[1]
	ones, _ := strconv.Atoi(os.Args[2])
	bits, _ := strconv.Atoi(os.Args[3])
	addr := net.ParseIP(dotAddr)
	if addr == nil {
		log.Default().Fatalln("Nil invalid address")
	}
	mask := net.CIDRMask(ones, bits)
	computedOnes, computedBits := mask.Size()
	network := addr.Mask(mask)
	fmt.Println("Address is", addr.String(),
		"\nMask length is", computedBits,
		"\nLeading Ones count is", computedOnes,
		"\nMask is (hex)", mask.String(),
		"\nNetwork is", network.String())
}
