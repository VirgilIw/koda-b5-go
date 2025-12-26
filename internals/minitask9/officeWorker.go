package minitask9

import (
	"fmt"
	"sync"
	"time"
)

func OfficeWorker() {
	var wg sync.WaitGroup
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	wg.Go(func() {
		Bath()
	})
	wg.Wait()
	wg.Go(func() {
		MakeCoffe()
	})
	wg.Wait()
	wg.Go(func() {
		Breakfast()
	})
	wg.Wait()
	wg.Go(func() {
		CleanBed()
	})
	wg.Wait()
	OtwKerja()
}

func Bath() {
	fmt.Println("mulai mandi dengan minyak sawit")
	time.Sleep(time.Second * 5)
	fmt.Println("mandi selesai, jangan lupa baby oil")
}

func MakeCoffe() {
	fmt.Println("kopi tubruk rasa kaki lima harga menurut golkar")
	time.Sleep(time.Second * 4)
	fmt.Println("selesai minum")
}

func Breakfast() {
	fmt.Println("makan bersama dia")
	time.Sleep(time.Second * 3)
	fmt.Println("makan enak finish")
}

func CleanBed() {
	fmt.Println("bersihkan tempat tidur pakai kuah pedes")
	time.Sleep(time.Second * 2)
	fmt.Println("tempat tidur rapih finish")
}

func OtwKerja() {
	fmt.Println("BERANGKAT KERJA")
}
