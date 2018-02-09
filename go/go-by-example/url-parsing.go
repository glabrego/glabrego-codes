package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	puts := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	puts(u.Scheme)

	puts(u.User)
	puts(u.User.Username())
	p, _ := u.User.Password()
	puts(p)

	puts(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	puts(host)
	puts(port)

	puts(u.Path)
	puts(u.Fragment)

	puts(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	puts(m)
	puts(m["k"][0])
}
