package minitask8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PaymentMethod() {
	fmt.Print("\nJawaban nomor 8: \n\n")
	var fiktifPayments []int
	// garbage collector, sampah dikumpulin
	scanner := bufio.NewScanner(os.Stdin)
	bayar := 20000

	for {
		questionPay()
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Printf("Silahkan bayar cash, total %d: ", bayar)

			scanner.Scan()
			jawaban1 := scanner.Text()

			uang, err := strconv.Atoi(jawaban1)
			if err != nil {
				fmt.Print("Baru mulai jangan aneh2 deh\n\n")
				break
			}

			for {
				if uang == bayar {
					fmt.Printf("Pembayaran dengan cash senilai Rp.%d berhasil.\n", uang)
					return
				} else if uang > bayar {
					kembalian := uang - bayar
					fmt.Printf("Pembayaran berhasil, ini kembalian anda Rp.%d.\n", kembalian)
					return
				} else {
					fmt.Println("Uang kamu kurang, yang bener aja?")
					fmt.Printf(`Kali ini masukin yang bener ya hensem 
⢿⣿⣿⣿⣧⡸⡎⡛⡩⠖⠀⣴⣿⣿⣿⠀⠀⠀⠀⠸⠇⠀⠙⢿⣿⣿⣿⣷⣌⢷⣑⢷⣄⠻⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⣫⠶⠛⠉⠀⠁⠀⠈⠈⠀⠠⠜⠻⣿⣆⢿⣼⣿⣿⣿
⢐⣿⣿⣿⣿⣧⢧⣧⢻⣦⢀⣹⣿⣿⣿⣇⠀⠄⠀⠀⠀⡀⠀⠈⢻⣿⣿⣿⣿⣷⣝⢦⡹⠷⡙⢿⣿⣿⣿⣿⣿⣿⣿⣿⠈⠁⠀⠀⠀⠁⠀⠀⠀⠱⣶⣄⡀⠀⠈⠛⠜⣿⣿⣿⣿
⠀⠊⢫⣿⣏⣿⡌⣼⣄⢫⡌⣿⣿⣿⣿⣿⣦⡈⠲⣄⣤⣤⡡⢀⣠⣿⣿⣿⣿⣿⣿⣷⣼⣍⢬⣦⡙⣿⣿⣿⣿⣿⣯⢁⡄⠀⡀⡀⠀⠄⢈⣠⢪⠀⣿⣿⣿⣦⠀⢉⢂⠹⡿⣿⣿
⠀⠀⠄⢹⢃⢻⣟⠙⣿⣦⠱⢻⣿⣿⣿⣿⣿⣿⣷⣬⣍⣭⣥⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⡙⢿⣼⡿⣿⣿⣿⣿⣿⣷⣄⠘⣱⢦⣤⡴⡿⢈⣼⣿⣿⣿⣇⣴⣶⣮⣅⢻⣿⡏		
: `)
					//
					scanner.Scan()
					tambahanStr := scanner.Text()
					tambahan, err := strconv.Atoi(tambahanStr)
					if err != nil {
						fmt.Println("woi isi yang bener dong!!")
						continue
					}
					uang = uang + tambahan
				}
			}
		case "2":
			fmt.Printf("Silahkan transfer, total Rp.%d: ", bayar)
			scanner.Scan()
			jawaban2 := scanner.Text()

			uang, err := strconv.Atoi(jawaban2)
			if err != nil {
				fmt.Print("\nsenyumku jadi pulsa :)\n\n")
				break
			}
			for {
				if uang == bayar {
					fmt.Printf("Transfer uang senilai Rp.%d berhasil.\n", uang)
					return
				} else if uang > bayar {
					kembalian := uang - bayar
					fmt.Printf("Pembayaran berhasil, ini kembalian anda Rp.%d.\n", kembalian)
					return
				} else {
					fmt.Println("Uang anda kurang, silahkan bayar sesuai ketentuan")
					fmt.Print("Masukkan uang tambahan: ")
					//
					scanner.Scan()
					tambahanStr := scanner.Text()
					tambahan, err := strconv.Atoi(tambahanStr)
					if err != nil {
						fmt.Println("1 kali it's ok, 2 kali it's ko")
						continue
					}
					uang = uang + tambahan
				}
			}
		case "3":
			fmt.Printf("Total pembayaran fiktif: Rp.%d, \nsilahkan bayar: ", bayar)
			scanner.Scan()
			input := scanner.Text()

			totalBayar, err := strconv.Atoi(input)
			if err != nil {
				fmt.Printf("\nsekali dua kali gapapa 2 kali 3 cari sendiri\n\n")
				break
			}

			if totalBayar > bayar {
				kembalian := totalBayar - bayar
				fiktifPayments = append(fiktifPayments, kembalian)
			} else if totalBayar < bayar {
				fmt.Println("Biaya pembayaran kurang")
				fmt.Print("tambah lagi uangnya: ")
				scanner.Scan()
				fiktifJawab := scanner.Text()
				fiktif, err := strconv.Atoi(fiktifJawab)
				if err != nil {
					fmt.Println("oke oke i'm fine")
				}
				fiktifPayments = append(fiktifPayments, fiktif+totalBayar)
			} else {
				fiktifPayments = append(fiktifPayments, totalBayar)
			}

		case "4":
			fmt.Println("exit")
			fmt.Println("Data pembayaran fiktif:", fiktifPayments)
			return

		default:
			fmt.Print("invalid choice PILIH YANG BENER\n\n")
		}
	}
}

func questionPay() {
	fmt.Println("Pilih metode pembayaran:")
	fmt.Println("1. Cash")
	fmt.Println("2. Transfer")
	fmt.Println("3. Fiktif")
	fmt.Println("4. Exit")
	fmt.Print("Masukkan pilihan: ")
}
