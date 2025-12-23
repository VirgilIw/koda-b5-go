package minitask2

import "fmt"

func Minitask2(limit int) {
	// cara for
	fmt.Println("Jawaban nomor 2:")
	for i := 1; i <= limit; i++ {
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("*")
	}

	// cara range
	// for i := range limit {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("*")
	// }
}
