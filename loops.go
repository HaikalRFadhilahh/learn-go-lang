/*
LOOPS FOR Merupakan teknik untuk melakukan perulangan dalam bahasa go-lang, hampir sama dengan bahasa lainnya untuk melakukan perulangan namun dalam golang hanya memiliki `for` untuk melakukan perulangan berbeda dengan bahasa pemrograman lainnya yang memiliki 3 atau lebih seperti While, While Do, dan For untuk membuat go-lang lebih sederhana maka golang hanya memiliki 1 teknik perulangan saya. Berikut tipe for dalam golang :
1. For Iteration (for i:= 0;i < 10;i++ {}) dengan menggunakan for ini kita bisa melakukan perulangan sesuai kondisi yang diinginkan seperti berapa jumlah perulangan yaitu 10 karena perulangan di mulai dari 0
2. For Range (for i,v := range array | slice | maps) dengan menggunakan range kita bisa melakukan perulangan untuk melakukan penampilan semua data yang ada dalam array / slice
3. Nested Loop dengan menggunakan teknik ini kita dapat menggunakan perulangan for di dalam for untuk membuat aplikasi kita dapat berjalan dengan lebih kompleks
Selain teknik tersebut ada 2 perintah penunjang dalam perulangan seperti :
1. break untuk melakukan terminate / keluar dari perulangan saat code break di eksekusi
2. continue untuk melakukan reset index dan perulangan di mulai dari atas ketika kdoe continou di eksekusi
! Gunakan perintah Loops sesuai dengan kondisi yang di butuhkan
*/
package main

import "fmt"

func main() {
	// Loops Iteration For
	for x := 5; x > 0; x-- {
		fmt.Printf("Ini Merupakan Perulangan Ke-%v \n", x)
	}

	// Loops Range (Harus Di Dukung Menggunakan Slice / Array / Maps)
	arrNamaMahasiswa := [...]string{"Haikal", "Zidan", "Vania", "Fajar"}
	for i, v := range arrNamaMahasiswa {
		fmt.Printf("[ Ini Adalah Index Ke-%v ]Halo, Selamat Pagi %v \n", i, v)
	}

	// Nested If
	for i := 0; i > 2; i++ {
		for x := 3; x > 0; x-- {
			fmt.Printf("Ini Adalah Perulangan Parent Ke-%v dan Perulangan Nested Ke-%v\n", i, x)
		}
	}

	// continue dan break Statement (Melakukan break Ketika Nama Daerah 'Gunung Kidul' dan Skip Ketika Nama 'Sleman')
	sliceNamaDaerah := []string{"Bantul", "Sleman", "Kota Yogyakarta", "Gunung Kidul", "Kulon Progo"}
	for _, v := range sliceNamaDaerah {
		// Skil Ketika Value Slice Sleman
		if v == "Sleman" {
			continue
		}

		// Lakukan Break Ketika Nama Daerah "Gunung Kidul"
		if v == "Gunung Kidul" {
			break
		}

		fmt.Printf("Selamat Datang di Yogyakarta Bagian %v\n", v)
	}
}
