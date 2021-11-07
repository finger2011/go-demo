package ninthweekserver

import (
	"bufio"
	"fmt"
	"net"
)

func ReceiveMessageByLength(length int) {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed err:", err)
			continue
		}
		go acceptData(conn, length)
	}
}

func acceptData(con net.Conn, length int) {
	defer con.Close()
	for {
		reader := bufio.NewReader(con)
		var buf = make([]byte, length * 10)
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("receive data:", recv)
	}
}
