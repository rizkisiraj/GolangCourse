package main

import (
	"fmt"
	"reflect"
)

func main() {
	var angka int = 42

	// Menggunakan reflect.ValueOf() untuk mendapatkan nilai dari variabel angka
	value := reflect.ValueOf(angka)

	// Menggunakan reflect.TypeOf() untuk mendapatkan tipe data dari variabel angka
	tipe := reflect.TypeOf(angka)

	fmt.Println("Nilai:", value)
	fmt.Println("Tipe data:", tipe)
}
