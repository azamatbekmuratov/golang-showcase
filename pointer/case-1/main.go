// pointer with a simple struct
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return fmt.Sprintf("User: [Name: %s, Age: %d]", u.Name, u.Age)
}

// pointer receiver
func (u *User) SetAge(age int) {
	u.Age = age
	fmt.Println(u)
}

func main() {
	u := User{
		Name: "John",
		Age:  25,
	}
	u.SetAge(30)
	fmt.Println(u)
}
