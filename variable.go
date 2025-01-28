/*
	Variable Golang adalah static ada beberapa tipe data seperti :
	int / int32 / int64 = Untuk Menyimpan Data Integer / Angka non koma
	float 32 / float64 = Untuk Menyimpan angka Desimal / koma
	string = Untuk menyimpan Data Kata / Kalimat / Huruf
	bool = Untuk Menyimpan Binary / True - False Data
*/

/*
	Deklarasi Tipe Data Ada 2 Cara Yaitu :
	1. `var nama_variabel tipe_data ` Dengan cara ini deklarasi variabel tidak harus langsung di isi namun harus eksplisit menentukan tipe data. Dan ketika tidak di isi setiap tipe data memiliki default value nya
	2. `nama_variabel := nilai / value` dengan cara ini lebih praktis tidak perlu menentukan tipe data namun tipe data di pasang sesuai data yang di masukkan
	! Pakai Cara Sesuai Dengan kebutuhan
*/

/*
	Cara Penulisan Variabel Untuk Konsistensi :
	Camel Case (namaMahasiswaLulus) Menggunakan Huruf Kapital Saat Kata Kedua dari Nama Variabel (Cocok Untuk Local Variable)
	Pascal Case (NamaMahasiswaLulus) Menggunakan Huruf Kapital Setiap Kata (Cocok Untuk Variabel Lintas Package)
	Snake Case (nama_mahasiswa_lulus) Menggunakan _ sebagai penghubung
	! Pastikan Untuk Konsistensi dalam penggunaan Cara Penulisan Variable
*/

package main

import "fmt"

func main() {
	// Deklarasi Tipe Data Langsung Intial Value / Langsung diisi (Ketika Variable Tidak Di pakai akan menyebabkan error)
	var namaMahasiswa string = "Haikal" // Harus Menentukan secara eksplisit
	namaOrganisasi := "Forum Asisten"   // Tipe Data Akan di asign sesuai value yaitu string

	// Menampilkan Variabel
	fmt.Println("Deklarasi Manual / Eksplisit :", namaMahasiswa)
	fmt.Println("Deklarasi Otomatis / By Value :", namaOrganisasi)

	// Deklarasi Value Kosong Dengan Eksplisit
	var umur int
	// umur := nil // Tidak Bisa Menggunakan Cara Seperti Ini karena Asign Otomatis tidak bisa default value

	// Lihat Value Default
	fmt.Println("Nilai Default Dari Variabel Umur / Int :", umur)

	// Multiple Declaration Value
	// var siswa1, siswa2 string // Ini Ketika Deklarasi 2 Variabel Langsung dengan Nilai Default
	var siswa1, siswa2 = "Haikal", "Fajar" // Ketika Langsung di asign
	// Atau
	// ! dosen1, dosen2 := "Andri", "Febri"
	// Atau
	var (
		nama      = "Haikal"
		totalAdek = 5
	)

	// Print All Variabel
	fmt.Println("Siswa 1 :", siswa1)
	fmt.Println("Siswa 2 :", siswa2)
	fmt.Println("nama :", nama)
	fmt.Println("Total Adek :", totalAdek)

}
