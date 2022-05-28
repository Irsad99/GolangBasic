package main

import (
	"basicgo/task2"
	"basicgo/task3"
	"fmt"
	"math"
)

func main() {

	// Number 1
	angka := 40.378
	fmt.Printf("Hasil Pembulatan %.2f",Round(angka))

	// Number 2
	task2.Show()

	// Number 3
	Balok := task3.Balok{Panjang: 7.3, Lebar: 5.3, Tinggi: 6.2}

	hitungBalok1(&Balok)

	Balok2 := &Balok
	Balok2.Panjang = 8.3
	hitungBalok2(Balok2)

	Balok3 := &Balok
	Balok3.Tinggi = 7.3
	hitungBalok3(Balok3)
}

func Round(angka float64) float64 {
	return (math.Round(angka*10)/10)
}

func hitungBalok1(benda task3.Hitung)  {
	fmt.Println("Luas : ", benda.Luas())
	fmt.Println("Keliling : ", benda.Keliling())
	fmt.Println("Volume : ", benda.Volume())
}

func hitungBalok2(benda task3.Hitung2d)  {
	fmt.Println("Luas : ", benda.Luas())
	fmt.Println("Keliling : ", benda.Keliling())
}

func hitungBalok3(benda task3.Hitung3d)  {
	fmt.Println("Volume : ", benda.Volume())
}