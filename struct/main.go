package main

import (
	"fmt"
	"struct/management"
)

func main() {
	user := management.User{}
	user.ID = 1
	user.FirstName = "zelda"
	user.LastName = "Suhada"
	user.Email = "suhada@gmail.com"
	user.IsActive = true

	user2 := management.User{
		ID:        2,
		FirstName: "suhada",
		LastName:  "ade",
		Email:     "suhada@gmail.com",
		IsActive:  false,
	}

	result := user.Display()
	fmt.Println(user2.Display())
	fmt.Println(result)

	fmt.Println(user)
	fmt.Println(user2)
	users := []management.User{user, user2}

	group := management.Group{"Gamer", user, users, true}
	fmt.Println(group)

	group.Display()
}
