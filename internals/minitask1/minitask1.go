package minitask1

import (
	"fmt"
	"math"
)

func Minitask1(phi float64, jari float64) {
	luas := hitungLuas(phi, jari)
	keliling := kelilingLingkaran(phi, jari)
	fmt.Println("Jawaban nomor 1:")
	fmt.Printf("Luas Lingkaran 		: %.2f\n", luas)
	fmt.Printf("Keliling Lingkaran	: %.2f\n", keliling)
	// Printf mirip backtick output stdout
	// Sprintf output berupa string
}

func hitungLuas(pi float64, jari float64) float64 {
	return pi * math.Pow(jari, 2)
}

// K = 2πr
func kelilingLingkaran(pi float64, jari float64) float64 {
	return 2 * pi * jari
}
