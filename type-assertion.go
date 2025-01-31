/*
	TYPE ASSERTION dalam Go-Lang Merupakan teknik untuk memperjelas suatu value tipe data non-spesifik seperti any atau Interface Kosong (interface{}) terkadang untuk menghindari error dalam menerima sebuah input value kebanyakan developer menggunakan interface{} atau any namun dalam golang tipe data generic seperti itu tidak dapat di olah secara mentah seperti menggunakan operator bahkan mengakses value jika any di isi dengan tipe data struct. TYPE ASSERTION menyelesaikan tersebut dengan dapat mengubah tipe data / menegaskan tipe data any menjadi tipe data statis.

	! Hati-Hati Dalam Menggunakan TYPE ASSERTION Selalu Gunakan Error Handling Karene Saat Any Tidak Support Dengan Tipe Data ASSERTION akan Return Panic Dan Aplikasi Akan Crash
	! Walaupun ketika Variabel interface{} atau any di lihat tipe datanya menggunakan library reflect.TypeOf() akan tetap menujukkan hasil yang sesuai dengan tipe data value. NamunTipe Data Value tidak dapat di gunakan untuk perhitungan untuk INT dan Manipulasi String
*/

package main

import "fmt"

/*
	Deklarasi Struct Untuk Data Mahasiswa
*/

type biodataMahasiswa struct {
	nama string
	umur int
}

func main() {
	// Membuat Variabel interface{} String
	var namaMahasiswa interface{} = "Haikal Raditya Fadhilah"

	// Melakukan Konversi Any Ke String
	_, ok := namaMahasiswa.(string)
	if ok {
		fmt.Println("Process Merubah Any / Interface{} Ke String Berhasil (namaMahasiswa)")
	} else {
		fmt.Println("Process Merubah Any / Interface{} Ke String Gagal (namaMahasiswa)")
	}

	// Membuat Variabel any int
	var umurMahasiswa any = 21

	if _, ok := umurMahasiswa.(int); ok {
		fmt.Println("Proses Merubah Any / interface{} Ke Int Berhasil (umurMahasiswa)")
	} else {
		fmt.Println("Proses Merubah Any / interface{} Ke Int Gagal (umurMahasiswa)")
	}

	// Membuat Variabel Interface{} Struct Mahasiswa
	var dataMahasiswa interface{} = biodataMahasiswa{
		nama: "Haikal Raditya Fadhilah",
		umur: 20,
	}

	// Merubah Interface{} Struct ke Struct biodataMahasiswa Agar Data Dalam Interface Dapat di Akses
	if v, ok := dataMahasiswa.(biodataMahasiswa); ok {
		fmt.Println("Konversi variabel dataMahasiswa ke Struct Access biodataMahasiswa Berhasil!!")
		fmt.Printf("Haloo, Selamat Malam %v\n", v.nama)
		fmt.Printf("Umur Kamu Adalah %v Tahun\n", v.umur)
	} else {
		fmt.Println("Konversi Variabel dataMahasiswa ke Struct Access biodataMahasiswa Berhasil!!")
	}
}
