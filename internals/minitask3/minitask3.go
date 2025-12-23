package minitask3

import "fmt"

func Minitask3() {
	fmt.Println("Jawaban nomor 3:")
	sliceOfInteger := []int{50, 75, 66, 20, 32, 90}
	fmt.Println(sliceOfInteger)
	//
	potong := sliceOfInteger[2:3]
	potong = append(potong, 75)
	potong[0] = 50
	//
	potongBaru := make([]int, len(potong))
	//
	copy(potongBaru, potong) // tidak referensi

	potongBaru = append(potongBaru, 66, 88, 20, 32, 90)
	fmt.Println("===========")
	for i, v := range potongBaru {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}
