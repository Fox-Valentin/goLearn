package struct_examp

import "fmt"

type Animal struct {
	color string
}

type Dog struct {
	Animal
	Name string
	Age  int
}
type Cat struct {
	Animal
	Name string
	Age  int
}

func Getdog() {
	// var dog Dog
	// dog := Dog{Name: "name", Age: 12}
	dog := new(Dog)
	dog.Name = "name"
	dog.Age = 12
	dog.color = "yellow"
	fmt.Println(dog)
	dog.Run()
}

// func (dog *Dog) Run() {
// 	fmt.Println("dog run!")
// }
func (animal *Animal) Run() {
	fmt.Println("Animal eat!")
}

func (dog *Dog) Run() string {
	return "Dog Run"
}
func (dog *Dog) Eat() string {
	return "Dog Eat"
}

func (cat *Cat) Run() string {
	return "Cat Run"
}
func (cat *Cat) Eat() string {
	return "Cat Eat"
}
