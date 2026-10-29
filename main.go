package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	cfg, err := loadConfig("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := newServer(*cfg)
	if err := s.listen(); err != nil {
		fmt.Println(err)
	}
}
