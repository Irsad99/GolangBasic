package task3

type Balok struct {
	Panjang float64
	Lebar float64
	Tinggi float64
}

func (blk Balok) Luas() float64 {
	p := blk.Panjang
	l := blk.Lebar
	t := blk.Tinggi

	result := (2*p*l) + (2*p*t) + (2*l*t)

	return result
}

func (blk Balok) Keliling() float64 {
	p := blk.Panjang
	l := blk.Lebar
	t := blk.Tinggi

	result := 4 * (p+l+t)

	return result
}

func (blk Balok) Volume() float64 {
	p := blk.Panjang
	l := blk.Lebar
	t := blk.Tinggi

	result := p*l*t

	return result
}