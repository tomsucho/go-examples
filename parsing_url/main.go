package main

import (
	"fmt"
	"net/url"
)

func main() {
	p := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, _ := url.Parse(s)
	p(u.Scheme)
	p(u.User)
	p(u.Query())
	p(u.Path)
	p(u.Fragment)
	p(u.RawQuery)
}
