package main

import (
	"fmt"
	"os"
)

func main() {
	p := fmt.Println
	pf := fmt.Printf

	str := `This is a smaple text!
	La la la!`
	err := os.WriteFile("file.txt", []byte(str), 0644)

	check(err)

	if data, err := os.ReadFile("file.txt"); err == nil {
		pf("%s\n", string(data))
	} else {
		p(err)
	}

}
