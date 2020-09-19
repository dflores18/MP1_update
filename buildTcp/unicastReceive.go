package buildTcp

import "bufio"
import "fmt"
import "log"
import "net"
// only needed below for sample processing
import "time"

func buildTCP() {
	ln, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		t := time.Now()
		thisTime := t.Format(time.RFC1123)
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:"+ string(message) + "from client. " + "System time is " + thisTime )
		// sample process for string received
		// send new string back to client
		conn.Write([]byte(message + "\n"))
		return
	}
}
