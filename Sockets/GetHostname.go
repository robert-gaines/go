package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	host, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	addr, err := net.LookupIP(host)

	if err != nil {
		panic(err)
	}

	fmt.Println(host)
	fmt.Println(addr[0])
}
