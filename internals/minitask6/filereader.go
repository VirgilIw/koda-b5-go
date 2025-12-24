package minitask6

import (
	"fmt"
	"os"
)

func Minitask6(path string) (err error) {
	// Minitask6()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovery from panic: %v\n", r)
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	//
	defer file.Close()
	//
	data, err := os.ReadFile(path)
	if err != nil {
		panic("panicManja")
	}
	//
	_, err = os.Stdout.Write(data)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}
