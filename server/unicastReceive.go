package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

//this function handles the connections to multiple clients
func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		t := time.Now()
		thisTime := t.Format(time.RFC1123)
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(netData)
		if temp == "STOP" {
			break
		}
		//timestamp for when the message was received
		fmt.Printf("Message Received:"+netData+"from address "+"%s "+"System time is "+thisTime, c.RemoteAddr().String())
		c.Write([]byte(netData))
	}
	c.Close()
}

func main() {
	//opens a tcp server where the clients can dial in
	l, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c) //goroutine to handle multiple processes/clients
	}
}
