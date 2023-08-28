package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:9090")
	fmt.Println("conn:", conn)

	defer conn.Close()
	var word string
	fmt.Println("输入一个英文单词")
	fmt.Scan(&word)
	conn.Write([]byte(word))
	buffer := make([]byte, 1024)
	conn.Read(buffer)
	fmt.Println(string(buffer))
}
