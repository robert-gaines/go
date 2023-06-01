package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Print("[+] Enter a domain name-> ")
	var domain string
	fmt.Scanln(&domain)

	addrs, err := net.LookupIP(domain)

	if err != nil {
		fmt.Println("[!] Error ", err)
		os.Exit(1)
	}

	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
