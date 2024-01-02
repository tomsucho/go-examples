// write a unit test for the SendUserInput function from tcp-client.go file, simulate connection using net.Pipe
package main

import (
	"io"
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

	// read the data written to the connection until EOF
	data, _ := io.ReadAll(conn2)

	// check if the data written matches the expected output
	expectedOutput := "Hello, World!\n"
	if string(data) != expectedOutput {
		t.Errorf("Expected output %s, got %s", expectedOutput, string(data))
	}
}
