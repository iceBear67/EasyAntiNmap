package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func handleConnection(conn net.Conn) {
	ip := strings.Split(conn.RemoteAddr().String(), ":")[0]
	fmt.Println(ip)
	if ip != "127.0.0.1" {
		fmt.Println(exec.Command("/bin/bash", "-c", "iptables -I INPUT -s "+ip+" -j DROP").Run())
	}

	conn.Close()
}
func main() {
	fmt.Println("Listening on 0.0.0.0:23")
	ln, err := net.Listen("tcp", ":23")
	if err != nil {
		fmt.Println("Failed to bind Port")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Unknown Error")
		}
		go handleConnection(conn)
	}
}
