package minitask5

import (
	"bufio"
	"fmt"
	"os"

	task1 "github.com/virgilIw/koda-b5-go/internals/minitask1"
	task2 "github.com/virgilIw/koda-b5-go/internals/minitask2"
	task3 "github.com/virgilIw/koda-b5-go/internals/minitask3"
	task4 "github.com/virgilIw/koda-b5-go/internals/minitask4"
)

func question() bool {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()

	switch choice {
	case "1":
		task1.Minitask1(22.0/7.0, 10.0)
	case "2":
		task2.Minitask2(5)
	case "3":
		task3.Minitask3()
	case "4":
		task4.Minitask4()
	case "0":
		exit()
		return false
	default:
		fmt.Println("Invalid choice")
	}
	return true
}
