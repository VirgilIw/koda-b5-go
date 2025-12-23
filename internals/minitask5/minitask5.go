package minitask5

func Minitask5() {

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
