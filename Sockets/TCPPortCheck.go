package main

import (
	"fmt"
	"net"
)

func checkPort(address string, port string) bool {
	conn, err := net.Dial("tcp", address+":"+port)

	if err != nil {
		return false
	}

	conn.Close()
	return true
}

func main() {
	fmt.Print("[+] Enter the IP address-> ")
	var address string
	var port string
	fmt.Scanln(&address)
	fmt.Print("[+] Enter the port to be checked-> ")
	fmt.Scanln(&port)

	if checkPort(address, port) {
		fmt.Printf("[*] Port %s is open ", port)
	} else {
		fmt.Printf("Port %s is closed", port)
	}
}
