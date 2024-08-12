package main

import (
	test11 "test11/cameltosnakecase"

	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := []string{
		"CamelCase",
		"camelCase",
		"HelloWorld",
		"132",
		" ",
		"",
		"A",
		"abcs",
		"AbC",
		"AbCEf",
		"abcAree",
		"ahe1Abde",
		"tesTing1",
		"SOME_VARIABLE",
		"ASuperLonGVariableName",
		"thisIsaTestOfCamelCase",
		"aA",
	}
	for _, arg := range args {
		challenge.Function("CamelToSnakeCase", test11.CamelToSnakeCase, solutions.CamelToSnakeCase, arg)
	}
}
