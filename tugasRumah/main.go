package main

import "fmt"

type Kendaraan interface {
	JarakTempuh(bensin int) int
	GetNama() string
}

type Motor struct {
	Kecepatan int
	Nama      string
}

type Mobil struct {
	Kecepatan int
	Nama      string
}

type Bajaj struct {
	Kecepatan int
	Nama      string
}

func (m Motor) JarakTempuh(bensin int) int {
	return bensin * 3
}

func (m Motor) GetNama() string {
	return m.Nama
}

func (m Mobil) JarakTempuh(bensin int) int {
	return bensin * 1
}

func (m Mobil) GetNama() string {
	return m.Nama
}

func (b Bajaj) JarakTempuh(bensin int) int {
	return bensin * 4
}

func (b Bajaj) GetNama() string {
	return b.Nama
}

// Fungsi untuk menentukan kendaraan yang paling efisien
func KendaraanPalingEfisien(bensin int, kendaraan ...Kendaraan) {
	var kendaraanPalingEfisien Kendaraan
	maxJarak := 0

	for _, k := range kendaraan {
		jarak := k.JarakTempuh(bensin)
		fmt.Printf("%s dapat menempuh: %d km\n", k.GetNama(), jarak)

		if jarak > maxJarak {
			maxJarak = jarak
			kendaraanPalingEfisien = k
		}
	}

	// Menampilkan kendaraan paling efisien berdasarkan jarak tempuh
	fmt.Printf("\nKendaraan paling efisien adalah %s yang dapat menempuh jarak: %d km\n", kendaraanPalingEfisien.GetNama(), maxJarak)
}

func main() {
	// Deklarasi objek struct kendaraan dengan nama
	motor := Motor{Kecepatan: 100, Nama: "Motor"}
	mobil := Mobil{Kecepatan: 80, Nama: "Mobil"}
	bajaj := Bajaj{Kecepatan: 60, Nama: "Bajaj"}

	// Jumlah bensin 10 liter
	bensin := 20

	// Menentukan kendaraan paling efisien
	KendaraanPalingEfisien(bensin, motor, mobil, bajaj)
}
