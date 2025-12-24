package minitask4

import "fmt"

// Buat struct yang bisa digunakan untuk menampung variabel data diri
// Nama
// Foto
// Email
// Umur
// Nomor telepon
// Status pernikahan
// Riwayat pendidikan (memiliki 2 properti, yaitu nama dan jurusan)

func Minitask4() {
	fmt.Println("Jawaban nomor 4:")

	user1 := Biodata{
		Name:        "Virgil I Ambar",
		Photo:       "https://static.vecteezy.com/system/resources/previews/026/658/763/non_2x/santa-claus-isolated-photo.jpg",
		Email:       "test@gmail.com",
		Age:         25,
		PhoneNumber: "083452127891",
		IsMarried:   false,
		Education: []Education{
			{
				Name:  "Elementary School",
				Major: "",
			},
			{
				Name:  "Junior High School",
				Major: "",
			},
			{
				Name:  "Senior High School",
				Major: "Science",
			},
			{
				Name:  "Bachelor Degree",
				Major: "Urban and Regional Planning Engineer",
			},
		},
	}

	fmt.Printf("%v\n", user1)
}
