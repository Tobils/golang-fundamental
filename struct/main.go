package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	name        string
	admin       User
	users       []User
	IsAvailable bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "zelda"
	user.LastName = "Suhada"
	user.Email = "suhada@gmail.com"
	user.IsActive = true

	user2 := User{
		ID:        2,
		FirstName: "suhada",
		LastName:  "ade",
		Email:     "suhada@gmail.com",
		IsActive:  false,
	}

	fmt.Println(user)
	fmt.Println(user2)

	res := displayUser(user)
	fmt.Println(res)

	users := []User{user, user2}

	group := Group{"Gamer", user, users, true}
	fmt.Println(group)

	displayGroup(group)
}

func displayUser(user User) string {
	result := fmt.Sprintf("name: %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}

func displayGroup(group Group) {
	fmt.Printf("Name Group: %s", group.name)
	fmt.Println("")
	fmt.Printf("Member Group: %d", len(group.users))
	fmt.Println("")
	fmt.Println("Name :")

	for _, item := range group.users {
		fmt.Println(item.FirstName)
	}

}
