/*
	INTERFACE dalam Go-Lang merupakan sebuah kontrak / tipe data syarat untuk struct biasanya. Interface memiliki Definini Fungsi-Fungsi yang harus di miliki oleh sebuah struct method agar Sebuah Struct Dapat Berjalan Berdampingan Dengan Baik, dengan menggunakan Interface membuat kode yang kita buat menjadi lebih konsisten dan memiliki Fungsi yang seragam. Sehingga dalam melakukan pengembangan developer lebih cepat dalam mendeksi kesalahan kode dari awal.
*/

/*
	! Pada Dasarnya Interface Merupakan Sebuah Type Atau Abstraksi Yang Dapat Kita Gunakan Seperti int, string, float32,dll
	Namun dalam Go-Lang Memiliki kelebihan dalam Interface kita dapat mengirimkan Sebuah Struct Tanpa Mendeklarasikan Tipe Interfacenya. Secara Otomatis Golang Akan Mendeteksi Apakah Struct Yang Kita Beri Sesuai Dengan Syarat / Kontrak atau tidak
*/

/*
	! Studi Kasus
	Untuk mendemonstraskan sebuah Interface kita akan membuat Interface Bernama PerangkatListrik dan terdapat 2 Perangkat Listrik Yaitu Printer Dan Scanner nanti Dalam Kontrak Interface Kita Akan memberikan Syarat Perangkat Listrik Bahwa Setiap Perangkat Listrik Harus Memiliki Kegunaan Dan Total Listrik Yang Digunakan
*/

package main

import "fmt"

/*
	Membuat Interface perangkatListrik Untuk Mendeklarasikan Abstraksi fungsi yang harus ada dalam Struct (Jika Struktur Fungsi Tidak Sama  / Tidak Ada Maka Return Error)
*/

type perangkatListrik interface {
	fungsiPerangkat()        // Bisa Di Tambahkan Tipe Data Return Kalau Ada
	totalPenggunaanListrik() // Bisa Di Tambahkan Tipe Data Return Kalau Ada
}

/*
	Deklarasi Struct dan Method Printer Yang Memenuhi Syarat Dari Interface Perangkat Listrik
*/

type printer struct {
	merek           string
	warnaHasilPrint string
	watt            int
}

func (p printer) fungsiPerangkat() {
	fmt.Println("Printer Berfungsi Untuk Mencetak Dokumen Dari Digital Ke Fisik / Kertas")
}

func (p printer) totalPenggunaanListrik() {
	if p.watt > 500 {
		fmt.Printf("Printer Memakan Cukup Besar Daya Sebanyak %v\n", p.watt)
	} else {
		fmt.Printf("Printer Terlalu Banyak Menggunakan Daya Sebanyak %v\n", p.watt)
	}
}

/*
	Deklarasi Struct Dan Method Struct Scanner Yang Memenuhi Syarat Dari Interface Perangkat Listrik
*/

type scanner struct {
	merek  string
	ukuran string
	watt   int
}

func (s scanner) fungsiPerangkat() {
	fmt.Println("Scanner Berfungsi Untuk Mengubah Dokumen Fisik Menjadi Digital Dokumen")
}

func (s scanner) totalPenggunaanListrik() {
	if s.watt > 500 {
		fmt.Printf("Scanner Memakan Cukup Besar Daya Sebanyak %v\n", s.watt)
	} else {
		fmt.Printf("Scanner Terlalu Banyak Menggunakan Daya Sebanyak %v\n", s.watt)
	}
}

/*
	Membuat Function Untuk Memanggil Fungsi Dari Setiap Struct Yang Telah Memenuhi Interface
*/

func callStructMethod(pl perangkatListrik) {
	pl.fungsiPerangkat()
	pl.totalPenggunaanListrik()
}

/*
	Main Function
*/

/*
	Dalam Contoh Function Main Saat call Function 'callStructMethod' variable hpPrinter dan canonScanner dapat di masukkan kedalam fungsi callStructMethod dikarenakan mematuhi kontrak dari Interface. 2 2 nya juga dapat masuk walaupun tanpa di deklarasikan Sebagai interface perangkatListrik Secara Ekplisit Seperti hpPrinter namun karena memenuhi sehingga fungsi callStructMethod tetap menganggap bahwa variabel hpPrinter Sesuai Dengan Kontrak
*/

func main() {
	// Deklarasi Struct Dengan 2 Cara Yaitu Type Inference Dan Declaration By Interface
	// ! Inference
	hpPrinter := printer{
		merek:           "Hewlett-Packard",
		warnaHasilPrint: "Printer Warna",
		watt:            700,
	}
	// And

	// ! Declare By Interface
	var canonScanner perangkatListrik = scanner{
		merek:  "Canon",
		ukuran: "300 x 500",
		watt:   400,
	}

	// Memanggil Fungsi Dengan Paramater Yang membutuhkan Interface perangkatListrik
	callStructMethod(hpPrinter)
	callStructMethod(canonScanner)
}
