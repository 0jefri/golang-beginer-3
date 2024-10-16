package main

import "fmt"

type Value struct {
	nilai int
}

func tampilkanNilai(n *Value) int {
	n.nilai = 10
	return n.nilai
}

func main() {
	//ini variabke buat menampung isi nya
	value := Value{}

	result := tampilkanNilai(&value)

	fmt.Println("Ini bentuk tampilan Nilai dari fungsi: ", result)
}
