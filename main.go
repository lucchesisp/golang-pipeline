package main

import "fmt"

func main() {
	var version = "dev"

	fmt.Println(version)

	fmt.Println(hello())
}

func hello() string {
	return "Hello Glang"
}
