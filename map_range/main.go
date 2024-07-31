package main

import "fmt"

func main() {
	personly := map[string]int{
		"maria": 15000,
		"paul":  19000,
	}
	for _, char := range personly {

		fmt.Println(char)
	}
	if personly == nil {
		fmt.Println("map est null. on va creer une map.")
		personly = make(map[string]int)
	}
	fmt.Println(personly)
}
