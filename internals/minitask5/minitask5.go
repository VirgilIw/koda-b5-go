package minitask5

import "fmt"

func Minitask5() {
	fmt.Print("\nJAWABAN NO 5: \n\n")
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
