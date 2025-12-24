package minitask7

import "fmt"

type Person struct {
	Name        string
	Address     string
	PhoneNumber string
	// Greet	func
}

func NewPerson(name, address, phoneNumber string) *Person {
	return &Person{
		Name:        name,
		Address:     address,
		PhoneNumber: phoneNumber,
	}
}

func (p *Person) PrintPersonName() string {
	return fmt.Sprintf("Name: %s\nAddress: %s\nPhoneNumber: %s", p.Name, p.Address, p.PhoneNumber)
}

func (p *Person) Greet() string {
	return fmt.Sprintf("\nHallo %s", p.Name)
}

func (p *Person) SetChangePersonName(newName string) {
	p.Name = newName
}

// & = ampersand
// * = asterisk
