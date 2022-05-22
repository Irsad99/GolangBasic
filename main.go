package main

import (
	"basicgo/task2"
	"basicgo/task3"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Number 1
	Round(40.3786698685)

	// Number 2
	task2.Show()

	// Number 3
	Balok := task3.Balok{Panjang: 7.3, Lebar: 5.3, Tinggi: 6.2}

	hitungBalok1(Balok)

	Balok2 := &Balok
	Balok2.Panjang = 8.3
	hitungBalok2(Balok2)

	Balok3 := &Balok
	Balok3.Tinggi = 7.3
	hitungBalok3(Balok3)
}

func Round(angka float64) {
	// getAngka and convert float to string
	getAngka := strconv.FormatFloat(angka, 'f', -1, 64)
	splitAngka := strings.Split(getAngka, "")

	// Indexing
	satuan := len(splitAngka) - 1
	puluhan := len(splitAngka) - 2

	// convert string to int
	getSatuan, _ := strconv.Atoi(splitAngka[satuan])
	getPuluhan, _ := strconv.Atoi(splitAngka[puluhan])

	// proses pembulatan

	if getSatuan >= 5 {
		getPuluhan += 1
		splitAngka[puluhan] = strconv.Itoa(getPuluhan)
		splitAngka[satuan] = "0"
	} else {
		splitAngka[satuan] = "0"
	}

	// proses cetak ulang

	for i := 0; i < len(splitAngka); i++ {
		fmt.Printf(splitAngka[i])
	}
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