/*
	RECOVER Go-Lang Merupakan sebuah fungsi bawaan dari golang untuk menangkap Panic dari Aplikasi Golang yang sedang berjalan, dengan recover memungkinkan kita untuk mencegah aplikasi kita agr tidak mati / crash. Agar Fungsi Recover Dapat di gunakan dengan optimal biasanya programmer memanfaatkan defer sehingga ketika crash dapat di tangkap oleh Recover dan continue aplikasi.

*/

package main

import "fmt"

/*
	Contoh Penggunaan Recover Tanpa Defer (Jika Menggunakan Perintah Recover Tanpa Defer seperti kurang optimal dan bermafaat, Maka dalam contoh berikutnya kita akan mengkombinasikan Defer dan Recover sehingga dapat bekerja dengan optimal)
	! Saat Menggunakan Recover akan return value berupa Any yang dapat disimpan divariabel
*/

func main() {
	// Cara Menggunakan Recover Dan Menyimpan Ke Variabel
	r := recover()
	fmt.Printf("Isi Variabel 'r' dari Return any recover : %v \n", r)

	// Cara Menggunakan Recover Untuk melakukan Pengecekan Apakah System Panic atau Tidak
	if err := recover(); err != nil {
		fmt.Println("Ada Error Dari Panic, Isi Panic : ", err)
	} else {
		fmt.Println("Tidak Ada Panic / Tidak Ada Nilai Dari Fungsi Recover!")
	}
}
