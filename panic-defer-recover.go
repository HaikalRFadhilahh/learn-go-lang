/*
PANIC-DEFER-RECOVER dalam Go-Lang Merupakan sebuah fungsi yang menjadi kesatuan untuk menghandle sebuah Panic dalam golang untuk mengatasi sebuah Aplikasi Golang Berhenti Tanpa Kesalahan Penulis Kode, Ada kala aplikasi yang sedang berjalan berhenti tanpa developer ketahui, maka setiap aplikasi golang diperlukan sebuah fungsi yang dapat melakukan Handle Panic dengan benar. Dalam Kasus Ini Kita Akan Mengetahui bagaimana cara mensimulasikan sebuah Panic Dalam Golang dan Menanfaatkan Defer & Function & Recover Untuk Membuat Aplikasi Optimal
*/

/*
! Tips : Sebaiknya Defer Di Taruh Pada Fungsi Dilain Main Karena Fungsi Defer itu Melakukan Terminasi Pada Function, Ketika DiTarus Pada Fungsi Utama
*/
package main

import (
	"fmt"
)

/*
FUNCTION HANDLE ERROR dalam kode dibawah merupakan sebuah Handler Untuk Melakukan Recover Sehingga Aplikasi Bisa Bejalan Seperti Berikutnya.
*/
func handlePanic() {
	// Gunakan Recover Untuk Menangkap Dan Healing Agar Aplikasi Golang Dapat Berjalan Dengan Baik
	if r := recover(); r != nil {
		fmt.Printf("Aplikasi Go-Lang Mengalami Panic, Error Panic : %v\n", r)

	} else {
		fmt.Println("Aplikasi Berjalan Dengan Baik, Tidak Ada Error!")
	}

}

/*
Membuat Fungsi Untuk Mensimulasi Panic Error!
*/
func prosesPanic() {
	// Fungsi Defer Untuk Handling Error
	defer handlePanic()

	/* Kode Utama / Simulasi Panic */
	// Membuat Slice
	sliceNamaMahasiswa := []string{"Haikal", "Zidan", "Faiq", "Sigit", "Lusi", "Zahra"}

	// Melakukan Perulangan
	for _, v := range sliceNamaMahasiswa {
		if v != "Lusi" {
			fmt.Println("OK...")
			continue
		}

		panic("Aplikasi Terhenti, Ada Kesalahan Tak Terduga!")
	}

}

/*
Agar Defer Handle Error Dalam Berjalan Dengan Baik Perintah Defer dijalankan saat fungsi main di jalankan, Jika Karena Perintah Defer Belum Berjalan Maka Panic Akan Tetap Terjadi Karena Function Yang Akan Di Jalankan Defer Masih Belum Terkonfigurasi.
*/
func main() {
	// Memanggil Panic Process Yang Di Dalam Sudah Ada Fungsi Defer Handle Panic. Jadi Ketika Ada Error Pada Panic Dia Akan Terminate fungsi prosesPanic namun tetap menjalankan perintah berikutnya di Fungsi Parent / Main
	prosesPanic()

	fmt.Println("Aplikasi Selesai Sesuai Dengan Code Akhir Program!")
}
