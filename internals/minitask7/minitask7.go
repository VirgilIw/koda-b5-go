package minitask7

import "fmt"

func Minitask7() {
	fmt.Print("\nJawaban nomor 7: \n\n")
	person1 := Person{
		Name:        "lil claus",
		Address:     "Jl. menuju surga",
		PhoneNumber: "0861764773",
	}
	newPerson := NewPerson("virgil", "bogor", "0821323123")
	fmt.Println(person1.PrintPersonName())
	fmt.Println(newPerson.PrintPersonName())

	person1.SetChangePersonName("Shawty")
	fmt.Println(person1.Greet())
}
