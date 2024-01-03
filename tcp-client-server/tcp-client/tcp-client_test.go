// write a unit test for the SendUserInput function from tcp-client.go file, simulate connection using net.Pipe
package main

import (
	"log"
	"net"
	"strings"
	"testing"
)

func TestSendUserInput(t *testing.T) {
	// create two ends of connection using net.Pipe
	conn1, conn2 := net.Pipe()

	// Simulate user input to SendUserInput from linux stdin
	testString := "Hello, World!\n"

	// call SendUserInput function with the connection
	go SendUserInput(conn1, strings.NewReader(testString))

	// read the 1024 bytes written to the connection
	data := make([]byte, 1024)
	n, err := conn2.Read(data)
	if err != nil {
		log.Fatal("Error reading connection: ", err)
	}
	r_msg := string(data[:n])

	// check if the data written matches the expected output
	if r_msg != testString {
		t.Errorf("Expected output %s, got %s", testString, r_msg)
	}
}
