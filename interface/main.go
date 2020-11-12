package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	bangunDatar := PersegiPanjang{Panjang: 6, Lebar: 5}
	luas := seberapaLuas(bangunDatar)
	fmt.Println(luas)

	persegi := Persegi{Sisi: 4}
	luas = seberapaLuas(persegi)
	fmt.Println(luas)
}

func seberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
