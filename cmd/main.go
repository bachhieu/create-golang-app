package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: print <message>")
		return
	}

	message := strings.Join(os.Args[1:], " ")
	a := strings.Split(message, " ")
	if a[len(a)-1] == "help" {
		fmt.Println("bạn muốn giúp gì? Vui lòng nhấn sử dụng: example help")
		return
	}
	for _, i := range a {
		fmt.Println(string(i))
	}

	fmt.Println("done")

}
