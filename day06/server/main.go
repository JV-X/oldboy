package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println(err)
	}
	defer listen.Close()
	fmt.Println("server is waiting...")
	conn, err := listen.Accept()
	fmt.Println("conn is ", conn)
	data := make([]byte, 1024)
	n, _ := conn.Read(data)
	fmt.Println("data is ", data[:n])
	conn.Write([]byte("hello world"))
}
