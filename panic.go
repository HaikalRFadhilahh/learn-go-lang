/*
PANIC Go-Lang adalah perintah untuk menghentikan secara paksa sebuah Aplikasi / Program yang sedang berjalan. Dalam Kondisi apapun Panic dieksekusi akan berhentu saat itu. Panic juga menjadi default perintah ketika ada Kode atau Kesalahan tak terduga dalam aplikasi entah itu Tidak Dapat Terkoneksi, Typo Kode dan lain Sebagainya. Panic dapat di selamatkan atau di tunda dengan menggunakan Recover namun harus di dukung oleh defer maka ketika kode panic Defer yang berisi Script Recover akan menyelamatkan aplikasi dari Crash.
*/
package main

import "fmt"

/*
Contoh Kasus : Akan Ada Perulangan untuk membaca Slice Dengan For ketika data error maka akan return panic walaupun perulangan Belum Selesai.
*/
func main() {
	myMap := []string{"Haikal", "Fajar", "Vania", "Ariza", "Zahra", "Lusi", "Rais"}
	for _, v := range myMap {
		if v == "Ariza" {
			panic("Ada Kesalahan Data!")
		}
		fmt.Printf("Haloo %v, Selamat Datang\n", v)
	}
}

/*
	Berbeda dengan Break walaupu terkesan sama bahwa perulangan tidak selesai perintah Break tidak menghentikan seluruh proses aplikasi, berbeda dengan Panic yang akan Terminate Seluruh Aplikasi sehingga Aplikasi Akan Mati dan tidak dapat di gunakan lagi.
	! Berhati-Hati Lah Dengan Panic. Handle Panic Dengan Benar. Gunakan Perintah Defer Dan Recover untuk menghindari Aplikasi Crash
*/
