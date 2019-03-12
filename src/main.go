package main

import (
	"client"
	"fmt"
	"server"
)

func main() {
	pro := client.Cmd2Protocol("set apple 12")
	args, count := server.Protocol2Args(pro)
	fmt.Printf("args: %v\ncount: %d\n", args, count)
}
