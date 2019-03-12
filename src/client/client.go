package client

import (
	"fmt"
	"strings"
)

func Cmd2Protocol(cmd string) string {
	var proto string
	args := strings.Split(cmd, " ")
	for i, arg := range args {
		if i == 0 {
			proto = fmt.Sprintf("*%d\r\n", len(args))
		}
		proto += fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg)
	}
	return proto
}
