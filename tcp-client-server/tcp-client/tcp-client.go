package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Function that takes user input from stdin through linux pipe and sends it to the server
func SendUserInput(conn net.Conn, r io.Reader) {
	// Read input from stdin
	input := bufio.NewScanner(r)
	for input.Scan() {
		// Send input to server
		conn.Write([]byte(input.Text() + "\n"))
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	serverHostname := flag.String("hostname", "localhost", "Server hostname or IP")
	serverPort := flag.Int("port", 8080, "Server port to connect to")
	// parse "-" flag to pipe input from stdin to server instead of prompting user for input
	pipeInput := flag.Bool("pipe", false, "Pipe input from stdin to server")
	flag.Usage = func() {
		fmt.Println("Usage: tcp-client [options]")
		fmt.Println("Options:")
		flag.PrintDefaults()
	}
	flag.Parse()

	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", *serverHostname, *serverPort))
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if *pipeInput {
		SendUserInput(conn, os.Stdin)
	} else {
		msg := "Hello, World!\n"
		conn.Write([]byte(msg))
		log.Printf("Sent: %s\n", msg)
	}

	log.Println("Waiting for reply..")
	resp := make([]byte, 1024)
	n, err := conn.Read(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server response: %s\n", resp[:n])
}
