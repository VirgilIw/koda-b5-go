package minitask5

import "fmt"

func Minitask5() {
	fmt.Println("Jawaban no 5: ")
	for {
		showMenu()
		/* question() menghasilkan false,
		karna question bertipe bool :)
		*/
		if !question() {
			return
		}
	}
}
