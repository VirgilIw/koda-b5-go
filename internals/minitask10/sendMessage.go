package minitask10

import "sync"

// kirim pesan ke channel
func SendMessage(wg *sync.WaitGroup, name, message string, ch chan Message) {
	defer wg.Done()

	ch <- Message{
		Name:    name,
		Message: message,
	}
}
