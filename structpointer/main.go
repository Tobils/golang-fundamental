package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

func (student *Student) Graduate() {
	student.Name = student.Name + " S.Si"
}

func main() {
	student := Student{1, "ade suhada", 3.71}
	fmt.Println(student.Name)
	graduate(&student)
	fmt.Println(student.Name)
	student.Graduate()
	fmt.Println(student.Name)

}

// melewatkan parameter pointer
func graduate(student *Student) {
	student.Name = student.Name + " S.Si"
}
