package main

import (
	"fmt"
	"os"
)

type Person struct {
	id        string
	name      string
	alamat    string
	pekerjaan string
	alasan    string
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Tolong masukkan nama atau nomor absen")
		fmt.Println("Contoh: 'go run main.go Siahaan' atau 'go run main.go 2'")
		return
	}
	persons := initSlice()
	var persona Person

	for _, person := range persons {
		if person.id == os.Args[1] || person.name == os.Args[1] {
			persona = person
			break
		}
	}

	if persona.name != "" {
		fmt.Printf("ID: %s\n", persona.id)
		fmt.Printf("Nama: %s\n", persona.name)
		fmt.Printf("Alamat: %s\n", persona.alamat)
		fmt.Printf("Pekerjaan: %s\n", persona.pekerjaan)
		fmt.Printf("Alasan: %s\n", persona.alasan)
	} else {
		fmt.Println("Data tidak ada")
	}

}

func initSlice() []Person {
	persons := []Person{
		{
			id:        "1",
			name:      "Siahaan",
			alamat:    "Perumahan Pesona",
			pekerjaan: "Backend Developer",
			alasan:    "Yo ndak tau",
		},
		{
			id:        "2",
			name:      "Prima",
			alamat:    "Perumahan Gelora",
			pekerjaan: "Frontend Developer",
			alasan:    "Wong saya suka og",
		},
		{
			id:        "3",
			name:      "Sijorni",
			alamat:    "Perumahan Indah",
			pekerjaan: "Fullstack Developer",
			alasan:    "Sorry yee",
		},
	}

	return persons
}
