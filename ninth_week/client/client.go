package ninthweekclient

import (
	"fmt"
	"net"
)

//SendMessageByLength fix length
func SendMessageByLength(length int) {
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("client connect err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("clinet connect success")
	for i := 0; i < 10; i++ {
		message := "send messagge by length"
		var buf = make([]byte, length)
		for i := 0; i < len(buf); i++ {
			if i < len(message) {
				buf[i] = message[i]
			} else {
				buf[i] = ' '
			}
		}
		_, err := conn.Write(buf)
		if err != nil {
			fmt.Println("client send message err:", err)
		} else {
			fmt.Println("client send message success:", string(buf))
		}
	}
}
