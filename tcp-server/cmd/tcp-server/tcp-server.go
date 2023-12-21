package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"time"
)

func handleConnection(conn net.Conn, id int) {

	defer conn.Close()
	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	log.Printf("New connection id:%d from %s \n", id, conn.RemoteAddr())

	for {
		if text, err := tp.ReadLine(); err != nil {
			log.Printf("Connection id: %d got %s from %s closed!\n", id, err, conn.RemoteAddr())
			return
		} else {
			log.Print("id: ", id, " Received: ", text)
			// write back to client
			date := time.Now().Format(time.RFC3339Nano)
			if _, err := conn.Write([]byte(fmt.Sprintf("%-40s ==> SERVER ACKed!\n", date))); err != nil {
				log.Printf("Connection id: %d write back to %s failed!\n", id, conn.RemoteAddr())
			}
		}
	}
}

func main() {
	portNum := flag.Int("port", 8080, "Port number to listen on")
	flag.Parse()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *portNum))
	if err != nil {
		log.Fatal(err)
	}
	i := 1
	log.Printf("Listening on %d", *portNum)
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		file, err := conn.(*net.TCPConn).File()
		defer file.Close()
		go handleConnection(conn, int(file.Fd()))
		i++
	}
}
