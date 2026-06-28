package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	u := User{ID: 1, Name: "Alice"}
	fmt.Printf("User: %+v\n", u)
}
