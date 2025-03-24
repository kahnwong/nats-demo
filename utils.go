package main

import (
	"strconv"
	"strings"
)

func stringToFloat(s string) (int, error) {
	vInt, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return 0, err
	} else {
		return int(vInt), nil
	}
}
