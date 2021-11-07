package main

import (
	ninthweekclient "go-demo/ninth_week/client"
	ninthweekserver "go-demo/ninth_week/server"
	"time"
)

func main() {
	socketByLength()
	select{}
}

func socketByLength() {
	var length = 23
	go ninthweekserver.ReceiveMessageByLength(length)
	time.Sleep(1 * time.Second)
	go ninthweekclient.SendMessageByLength(length)
}
