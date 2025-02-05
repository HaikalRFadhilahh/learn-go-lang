/*
METHOD merupakan sebuah fungsi namun dalma implementasikannya dia diharuskan untuk mendapatkan instase cari sebuah type. Sehingga ketika sebuah variabel dideklarasikan menggunakan type dan ada method yang diimplementasikan maka tetap bisa memenuhi sebuah Imnterface.
! Method Tidak harus menggunakan Struct Namun yang Menjadi Patokan Adalah Type
! Implementasi Interface Tidak Harus Selalu Menggunakan Struct / Tetapi Bisa Selain Struct Yang di bungkus dalam type
*/

/*
! IMPORTANT
Saat mendeklarasikan Sebuah Type sama saja membuat tipe data baru bukan alias. Jadi Ketika semisal ada variabel string murni dan namaMahasiswa yaitu string dan ingin saling si masukkan harus di konversi terlebih dahulu dengan nama_tipe_data/type(variabel_beda_tipe_data)
*/
package main

import "fmt"

/*
	Membuat Interface
*/

type mhs interface {
	tambahGelar(g string) *namaMahasiswa
}

/*
	Membuat Type String
*/

type namaMahasiswa string

/*
Membuat Method Untuk Memenuhi Interface
*/
func (n *namaMahasiswa) tambahGelar(g string) *namaMahasiswa {
	*n = namaMahasiswa(fmt.Sprintf("%v, %v", *n, g))
	return n
}

func (n *namaMahasiswa) show() {
	fmt.Println(*n)
}

func main() {
	// Deklarasi Nama Variabel Dengan Interface
	var haik namaMahasiswa = "Haikal Raditya Fadhilah"
	var haikal mhs = &haik

	haikal.tambahGelar("S.Kom").show()
}
