package main

// minitask 1

import (
	"fmt"

	task1 "github.com/virgilIw/koda-b5-go/internals/minitask1"
	task2 "github.com/virgilIw/koda-b5-go/internals/minitask2"
	task3 "github.com/virgilIw/koda-b5-go/internals/minitask3"
	task4 "github.com/virgilIw/koda-b5-go/internals/minitask4"
	task5 "github.com/virgilIw/koda-b5-go/internals/minitask5"
)

func main() {
	var phi float64
	var jari float64
	//
	phi = 22 / 7
	jari = 10
	//
	limit := 5
	//
	task1.Minitask1(phi, jari)
	fmt.Println("===========")
	task2.Minitask2(limit)
	fmt.Println("===========")
	task3.Minitask3()
	fmt.Println("===========")
	task4.Minitask4()
	task5.Minitask5()
}
