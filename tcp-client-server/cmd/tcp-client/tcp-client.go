package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	serverHostname := flag.String("hostname", "localhost", "Server hostname or IP")
	serverPort := flag.Int("port", 8080, "Server port to connect to")

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", *serverHostname, *serverPort))
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	msg := "Hello, World!\n"
	conn.Write([]byte(msg))
	log.Printf("Sent: %s\n", msg)

	resp := make([]byte, 1024)
	n, err := conn.Read(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server response: %s\n", resp[:n])
}
