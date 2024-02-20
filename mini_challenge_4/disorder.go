package main

import (
	"fmt"
	"sync"
)

type latihan interface {
	munculNama() string
}

type Coba struct {
	nama string
}

type Bisa struct {
	nama string
}

func (b Bisa) munculNama() string {
	return b.nama
}

func (c Coba) munculNama() string {
	return c.nama
}

func cetak(listBarang []string, index int) {
	fmt.Println(listBarang, index)
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	arrayPertama := [3]string{
		Coba{nama: "coba1"}.munculNama(),
		Coba{nama: "coba2"}.munculNama(),
		Coba{nama: "coba3"}.munculNama(),
	}

	arrayKedua := [3]string{
		Bisa{nama: "bisa1"}.munculNama(),
		Bisa{nama: "bisa2"}.munculNama(),
		Bisa{nama: "bisa3"}.munculNama(),
	}

	for i := 1; i <= 4; i++ {
		go cetak(arrayPertama, i)
		go cetak(arrayKedua, i)
	}

}

// func main() {
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	for i := 1; i <= 4; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			mu.Lock()
// 			// Critical section starts
// 			fmt.Printf("[coba1 coba2 coba3] %d\n", i)
// 			// Simulating some work
// 			time.Sleep(time.Millisecond * 100)
// 			// Critical section ends
// 			mu.Unlock()
// 		}(i)
// 	}
// 	wg.Wait() // Wait for all goroutines to finish
// }
