package main

import "fmt"

//buat struct dengan properti nama,score. buatkan fungsi untuk merubah score.

type Person struct {
	Name  string
	Score int
}

func ubahScore(s *Person, scoreBaru int) {
	s.Score = scoreBaru
}

func main() {
	newPerson := Person{
		Name:  "Jefri",
		Score: 85,
	}

	fmt.Println("Sebelum diubah:", newPerson)

	ubahScore(&newPerson, 100)

	fmt.Println("setelah diubah:", newPerson)
}
