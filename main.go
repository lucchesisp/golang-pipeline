package main

import "fmt"

func main() {
	var version = "dev"

	fmt.Println(version)

	fmt.Println(hello())
}

func hello() string {
	fmt.Println("teste ci")
  return "Hello Glang"
}
