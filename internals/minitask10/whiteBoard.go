package minitask10

import "fmt"

func WhiteBoard(ch chan Message) {
	for pesan := range ch {
		fmt.Printf("%s: %s\n", pesan.Name, pesan.Message)

	}
	fmt.Println("papan tulis ditutup")
}
