package main

type backend struct {
	address     string
	healthy     bool
	activeConns int
}
