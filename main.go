package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	s := newServer(":5555")
	if err := s.listen(); err != nil {
		fmt.Println(err)
	}
}
