package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	t := time.Now()
	p(t.Add(time.Hour * 24))
	then := t.AddDate(0, 1, 0)
	p(then)
	p(then.Sub(t))

	future := time.Date(2023, 7, 28, 14, 30, 20, 1000, time.Local)
	p(future)
}
