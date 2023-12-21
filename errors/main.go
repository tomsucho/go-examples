package main

import (
	"fmt"
)

type myError struct {
	id  int
	msg string
}

func (my_err myError) Error() string {
	return fmt.Sprintf("ERROR: id=%d, msg=%s", my_err.id, my_err.msg)
}

func divider(a, b int) (int, error) {
	if b == 0 {
		return 0, myError{id: 1, msg: "Can't divide by 0!"}
	} else {
		return a / b, nil
	}

}

func main() {

	fmt.Println("Hello World!")
	if v, err := divider(4, 0); err == nil {
		fmt.Printf("Result = %d", v)
	} else {
		fmt.Println(err)
	}

}
