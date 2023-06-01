package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	var network string

	fmt.Print("[+] Enter the network in CIDR notation-> ")
	fmt.Scanln(&network)

	_, ipnet, err := net.ParseCIDR(network)

	if err != nil {
		fmt.Println("[!] Error parsing CIDR:", err)
		return
	}

	mask := binary.BigEndian.Uint32(ipnet.Mask)
	start := binary.BigEndian.Uint32(ipnet.IP)

	finish := (start & mask) | (mask ^ 0xffffffff)

	// loop through addresses as uint32
	for i := start; i <= finish; i++ {
		// convert back to net.IP
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		fmt.Println(ip)
	}
}
