package task2

import "fmt"

func Show() {
	deret := deretBilangan{40}
	deret_bil := &deret

	// prima
	prima := deret_bil.prima()
	fmt.Println("Bil. Prima : ", prima)

	// genap
	genap := deret_bil.genap()
	fmt.Println("Bil. Genap : ", genap)

	deret_bil.limit = 60

	// ganjil
	ganjil := deret_bil.ganjil()
	fmt.Println("Bil. Ganjil : ", ganjil)

	// fibo
	fibo := deret_bil.fibo()
	fmt.Println("Bil. Fibonacci : ", fibo)
}

type deretBilangan struct {
	limit int
}

func (deret deretBilangan) prima() []int {
	result := []int{}
	for i := 1; i < deret.limit; i++ {
		fact := 0
		for j := 1; j < deret.limit; j++ {
			if i%j == 0 {
				fact++
			}
		}
		if fact == 2 && i != 1 {
			result = append(result, i)
		}
	}
	return result
}

func (deret deretBilangan) genap() []int {
	result := []int{}
	for i := 0; i < deret.limit; i++ {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func (deret deretBilangan) ganjil() []int {
	result := []int{}
	for i := 0; i < deret.limit; i++ {
		if i%2 != 0 {
			result = append(result, i)
		}
	}
	return result
}

func (deret deretBilangan) fibo() []int {
	result := []int{}
	bil1, bil2, bil3 := 0, 1, 0

	for i := 1; i < deret.limit; i++ {
		if bil1 < deret.limit {
			result = append(result, bil1)
		}
		bil3 = bil1 + bil2
		bil1 = bil2
		bil2 = bil3
	}
	return result
}