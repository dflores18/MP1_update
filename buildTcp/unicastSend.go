package buildTcp

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"
)

func buildClient() {

	addr, _ := net.ResolveTCPAddr("tcp", ":8081")
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		panic(err.Error())
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Send message: ")
	text, _ := reader.ReadString('\n')
	timeout := rand.Intn(5)
	select {
	case <- time.After(time.Duration(timeout) * time.Second):
		t := time.Now()
		thisTime := t.Format(time.RFC1123)
		fmt.Fprintf(conn, text+"\n")
		fmt.Print("Received " + text + "from server. " + "System time is " + thisTime + "\n")
		conn.Close()
	}
}


