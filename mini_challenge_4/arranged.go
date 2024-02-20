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

func cetak(wg *sync.WaitGroup, mutex *sync.Mutex, listBarang []string, index int) {
	defer wg.Done()
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
		wg.Add(2)
		mutex.Lock()
		go cetak(&wg, &mutex, arrayPertama[:], i)
		mutex.Unlock()
		mutex.Lock()
		go cetak(&wg, &mutex, arrayKedua[:], i)
		mutex.Unlock()
	}

	wg.Wait()

}
