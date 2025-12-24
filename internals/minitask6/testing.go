package minitask6

import (
	"fmt"
)

func Testing() {
	err1 := Minitask6("catatan.txt")
	if err1 != nil {
		fmt.Println("error:", err1)
		return
	}
	err2 := Minitask6("catatan.txt")
	if err2 != nil {
		fmt.Println("error:", err2)
		return
	}
	err3 := Minitask6("/home/Desktop/koda5_go/internals/minitask6/filereader.go")
	if err3 != nil {
		fmt.Println("error:", err3)
		return
	}
	fmt.Println("program selesai")
}
