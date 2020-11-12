package management

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

/**
fungsi ini dimiliki oleh instanse user
*/
func (user User) Display() string {
	result := fmt.Sprintf("name: %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) Display() {
	fmt.Printf("Name Group: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Group: %d", len(group.Users))
	fmt.Println("")
	fmt.Println("Name :")

	for _, item := range group.Users {
		fmt.Println(item.FirstName)
	}
}
