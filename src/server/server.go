package server

import (
	"strconv"
	"strings"
)

func Protocol2Args(protocol string) ([]string, int) {
	argv := make([]string, 0)
	argn := 0

	args := strings.Split(strings.Trim(protocol, " "), "\r\n")
	if len(args) == 0 {
		return nil, 0
	}
	argn, err := strconv.Atoi(args[0][1:])
	if err != nil {
		return nil, 0
	}

	var argsLen []int
	var j int
	for _, arg := range args {
		if len(arg) == 0 {
			continue
		}
		if arg[0] == '$' {
			alen, err := strconv.Atoi(arg[1:])
			if err == nil {
				argsLen = append(argsLen, alen)
			}
		} else {
			if j < len(argsLen) && argsLen[j] == len(arg) {
				argv = append(argv, arg)
				j++
			}
		}
	}
	return argv, argn
}
