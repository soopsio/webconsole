package main

import (
	"fmt"
	"net/url"
)

func main() {

	//golang 1.8 bug
	u, err := url.Parse("//192.168.2.2:6688")
	if nil != err {
		panic(err)
	}
	fmt.Println(u.Host, u.Port())
}
