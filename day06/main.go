package main

import (
	"fmt"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:9090")
	defer conn.Close()
	for true {
		var word string
		fmt.Println("请输入。。。")
		fmt.Scan(&word)
		conn.Write([]byte(word))
	}
}
