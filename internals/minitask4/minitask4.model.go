package minitask4
type Education struct {
	Name  string
	Major string
}
type Biodata struct {
	Name        string
	Photo       string
	Email       string
	Age         int
	PhoneNumber string
	IsMarried   bool
	Education   []Education
}

