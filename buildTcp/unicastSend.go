package buildTcp

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func Send(maxDelay int) {
	//Send() dials in to the server creating a client
	addr, _ := net.ResolveTCPAddr("tcp", ":8081")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err.Error()) //handles error
	}
	timeout := rand.Intn(maxDelay)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("send message: ")
	text, _ := reader.ReadString('\n')
	//implemented delay using cases and the After method
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Fprintf(conn, text+"\n")
		t := time.Now()
		thisTime := t.Format(time.RFC1123)
		text, _ := bufio.NewReader(conn).ReadString('\n')
		//sent timestamp of message
		fmt.Print("Sent " + text + "to server " + "127.0.0.1 " + "System time is " + thisTime + "\n")
		conn.Close()
	}
}
