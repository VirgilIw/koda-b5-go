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
	user1 := biodata{
		name:        "Virgil I Ambar",
		photo:       "https://static.vecteezy.com/system/resources/previews/026/658/763/non_2x/santa-claus-isolated-photo.jpg",
		email:       "test@gmail.com",
		age:         25,
		phoneNumber: "083452127891",
		isMarried:   false,
		education: []string{
			"Elementary School", "Junior High School",
			"senior High School", "Bachelor degree",
		},
	}
	fmt.Println(user1)
}
