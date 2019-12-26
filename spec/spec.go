package main

import "fmt"

// User - recs about user
type User struct {
	Name    string
	Account float32
}

func main() {
	usersAccount := map[int]User{}

	usersAccount[0] = User{Account: 300.25, Name: "Vasily"}
	usersAccount[1] = User{Account: 399.99, Name: "Bassy"}

	fmt.Println(usersAccount)

	//var rec User = usersAccount[1]
	//usersAccount[1] = rec

	var x = User{}
	x.Name = "Bess"
	usersAccount[1] = x

	fmt.Println(usersAccount)

}
