package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("input parameters: raw,simple,orm,update,insert")
		return
	}
	if os.Args[1] == "raw" {
		rawCommand()
	} else if os.Args[1] == "simple" {
		simpleCommand()
	}
}
