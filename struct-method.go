/*
	STRUCT METHOD Dalam Go-Lang Merupakan Sebuah Teknik Membuat Function namun hanya dapat di eksekusi atau digunakan dari instance Struct Yang Telah DI deklarasikan. Berbeda Dengan Function Yang Dapat Diekskusi Tanpa Instance / Turunan Dari Struct.
	! Kita Akan membuat struct yang aturannya Satu Orang Akan Memiliki 1 Rumah / Alamat. Setiap Struct Memiliki Struct Method
*/

package main

import "fmt"

/*
	Membuat Struct Orang
*/

type orang struct {
	nama  string
	umur  int
	rumah alamat
}

/*
	Membuat Struct Method Untuk Menampilkan Nama
*/

func (o orang) getNama() orang {
	fmt.Printf("Nama Saya Adalah %v \n", o.nama)
	return o
}

/*
Membuat Method Untuk Menampilkan Jenis Rumah Namun Dari Struct Orang (orang->rumah->tipe)
! Ketika Ingin membuat Chain Struct method Pastikan Untuk Return Struct Value Sebelum Next Method Di Eksekusi
*/
func (o orang) getTypeHome() {
	fmt.Printf("Saya Memiliki Tipe Rumah Yaitu %v\n", o.rumah.tipe)
}

/*
	Membuat Struct Alamat
*/

type alamat struct {
	tipe   string
	lokasi string
}

/*
	Membuat Struct Untuk Menampilkan Jenis Rumah
*/

func (a alamat) getTipeRumah() {
	fmt.Printf("Saya Memiliki Tipe Rumah Yaitu %v\n", a.tipe)
}

func main() {
	// Deklarasi Struct Child Dulu Yaitu Alamat
	rumahHaikal := alamat{
		tipe:   "Apartment",
		lokasi: "Kota Yogyakarta",
	}

	// Deklarasi Struct Parent Yaitu Orang
	haikal := orang{
		nama:  "Haikal Raditya Fadhilah",
		umur:  21,
		rumah: rumahHaikal,
	}

	// Panggil Struct Method Dari Get Nama
	haikal.getNama().getTypeHome()

	// Panggil Jenis Rumah getTipeRumah
	haikal.rumah.getTipeRumah()
}
