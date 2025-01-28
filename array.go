/*
	Array Merupakan Tipe Data Lanjutan untuk menyimpan beberapa data sekaligus namun masih dalam 1 tipe data, Array bersifat fixed artinya tidak dapat di rubah. ada beberapa cara mendeklarasikan array apakah hanya inisiasi atau langsung mengisi value bahkan menghitung jumlah data berdasarkan value yang di asign. Berikut Contoh Deklarasi Array :
	Deklarasi Array Kosong :
	- var nama_array = [jumlah_data_array]tipe_data{}
	atau
	- nama_array := [jumlah_data_array]tipe_data{}

	Deklarasi Array Setengah :
	- var nama_array = [10]tipe_data{1,2,3,4,5}
	atau
	- nama_array := [10]tipe_data{1,2,3,4,5}

	Deklarasi Jumlah Array Dengan Default Data :
	- var nama_array = [...]string{"Haikal","Zidan","Reza"}
	atau
	- nama_array := [...]string{"Haikal","Zidan","Reza"}
	! Gunakan Sesuai Kondisi
*/

package main

import "fmt"

func main() {
	// Deklarasi Array Kosong
	var arrNamaMahasiswa = [3]string{""}

	// Deklarasi Array Setengah Isi
	arrNamaDaerah := [5]string{"Bantul", "Sleman"}

	// Deklarasi Array By Value
	arrTahun := [...]int{2020, 2021, 2022, 2023, 2024}

	// Print Data
	fmt.Println("Data Kosong Array Mahasiswa :", arrNamaMahasiswa)
	fmt.Println("Ukuran Jumlah Data Array Nama Mahasiswa :", len(arrNamaMahasiswa))

	fmt.Println("Data Array Nama Daerah :", arrNamaDaerah)
	fmt.Println("Ukuran Jumlah Data Array Nama Daerah :", len(arrNamaDaerah))

	fmt.Println("Data Array Tahun :", arrTahun)
	fmt.Println("Ukuran Jumlah Dara Array Tahun :", len(arrTahun))
}
