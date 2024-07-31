package funnc_work

import (
	"errors"
	"strconv"
)

func Add(a, b string) (int, error) {
	n1, err := strconv.Atoi(a)
	if err != nil {
		return 0, errors.New("invalid first number")
	}
	n2, err := strconv.Atoi(b)
	if err != nil {
		return 0, errors.New("invalid second number")
	}
	return n1 + n2, nil
}
