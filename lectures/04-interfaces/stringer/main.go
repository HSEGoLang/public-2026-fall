package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("User(%d:%q)", u.ID, u.Name)
}

func main() {
	u := User{ID: 42, Name: "Ada"}
	fmt.Printf("%v\n", u)
	fmt.Printf("%s\n", u)
	fmt.Printf("%#v\n", u)
}
