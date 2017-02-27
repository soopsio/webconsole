package main

import (
	"apibox.club/utils"
	"fmt"
	"net"
	"net/url"
	"strings"
)

func main() {

	//golang 1.8 bug
	u, err := url.Parse("192.168.2.2:6688")
	if nil != err {
		panic(err)
	}
	fmt.Println(u.Host, u.Port())
}
