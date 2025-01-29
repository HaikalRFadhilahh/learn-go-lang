/*
	DEFER Pada GO-Lang Merupakan perintah untuk mengeksekusi suatu Kode atau Function pada saat aplikasi Berhenti Secara Grateful (Baik-Baik) atau Crash (Panic Error). DEFER memungkinkan developer untuk melakukan pembersihan Resouce dari Shutdown Database Connection, Menyelamatkan Aplikasi, Ataupun Reboot Aplikasi.
	! Walaupun Defer Lebih Dari 1 Namun Defer Yang Pertama DI Eksekusi adalah Defer Yang Di Deklarasi Paling Terakhir
*/

package main

import "fmt"

/*
Membuat Function yang akan di eksekusi Defer Sebelum Aplikasi Tertutup Atau Crash
*/
func crashHandler() {
	fmt.Println("Function Yang Di Eksekusi Oleh Defer Sebelum Aplikasi Mati!, Ini BLOCK KODE FUNCTION BIASA")
}

func main() {
	// Melakukan Defer Dari Anonymous Function
	defer func() {
		fmt.Println("Function Yang Di Eksekusi Defer Sebelum Selesai / Crash, Ini BLOCK KODE ANONYMOUS FUNCTION")
	}()

	// Defer Dengan Function Reguler
	defer crashHandler()

	// Defer Dengan 1 Syntax
	defer fmt.Println("Syntax FMT yang di eksekusi Defer Sebelum Selesai, INI SYNTAX DARI LIBRARY FMT")

	// Syntax Biasa Yang Di Eksekusi Terakhir Namun Muncul Paling Awal Karena Defer Akan Eksekusi Paling Terkahir
	fmt.Println("Hello Gaisss!")
}
