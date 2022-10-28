package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
		fmt.Fprintln(conn, "OK")
	}

}
func main() {
	//监听一个端口
	ln, _ := net.Listen("tcp", ":8030")
	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}

}
