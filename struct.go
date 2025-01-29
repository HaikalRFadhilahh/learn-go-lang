/*
	STRUCT dalam golang meruapakn sebuah tipe_data lanjutan yang dapat menampung berbagai variable lainnya sehingga terstruktur dan lebih mudah dalam memanajemen variabel yang bisa saja menjadi satu paket tanpa harus mendeklarasikan satu persatu sehingga memiliki kemungkinan untuk tertinggal untuk dideklarasikan. Struct meruapakan tipe data artinya bisa menjadi tipe data saat kita mendeklarasikan variabel dan memberikan paramater ke fungsi. Ada Dua cara pembuatan variabel struct yaitu :
	1. Deklarasi Kosong (var nama_variabel nama_struct) dengan ini anda dapat membuat struct dengan nilai default dari setiap element struct
	2. Deklarasi Secara Langsung Dengan Value (nama_variabel := nama_struct{...data}) dengan cara ini kita bisa melakukan deklarasi secara langsung dan asign nilai sehingga tidak perlu memasukkan saya satu persatu
*/

package main

import "fmt"

/*
Mendeklarasikan Struct Dosen Pembimbing
*/
type dosenPembimbing struct {
	namaDosen  string
	prodiDosen string
}

/*
Mendeklarasikan Struct Mahasiswa
*/
type mahasiswa struct {
	namaMahasiswa            string
	npmMahasiswa             string
	dosenPembimbingMahasiswa dosenPembimbing
}

/*
Deklarasi Fungsi Untuk Pemanggilan Struct dengan Memberikan paramater tipe structnya ke fungsi
*/
func printStruct(m mahasiswa) {
	// Pemanggilan Struct Haikal
	fmt.Print("Data Diri (Dipanggil Secara Otomatis Menggunakan Fungsi dan Paramater Struct)\n")
	fmt.Printf("Nama Mahasiswa : %v\n", m.namaMahasiswa)
	fmt.Printf("NPM Mahasiswa : %v\n", m.npmMahasiswa)
	fmt.Printf("\n")
	fmt.Printf("Data Diri Dosen Pembimbing \n")
	fmt.Printf("Dosen Pembimbing : %v\n", m.dosenPembimbingMahasiswa.namaDosen)
	fmt.Printf("Asal Prodi Dosen Pembimbing : %v\n", m.dosenPembimbingMahasiswa.prodiDosen)
	fmt.Printf("\n")
}

func main() {
	// Deklarasi Haikal Dan Dosen Pembimbing Haikal
	var haikal mahasiswa
	var dosenPembimbingHaikal dosenPembimbing

	// Melakukan Pengisian Data
	dosenPembimbingHaikal.namaDosen = "Pak Ika"
	dosenPembimbingHaikal.prodiDosen = "Sistem Informasi"

	haikal.namaMahasiswa = "Haikal Raditya Fadhilah"
	haikal.npmMahasiswa = "21.11.3910"
	haikal.dosenPembimbingMahasiswa = dosenPembimbingHaikal

	// Pemanggilan Struct Haikal
	fmt.Print("Data Diri (Dipanggil Secara Eksplisit / Manual)\n")
	fmt.Printf("Nama Mahasiswa : %v\n", haikal.namaMahasiswa)
	fmt.Printf("NPM Mahasiswa : %v\n", haikal.npmMahasiswa)
	fmt.Printf("\n")
	fmt.Printf("Data Diri Dosen Pembimbing \n")
	fmt.Printf("Dosen Pembimbing : %v\n", haikal.dosenPembimbingMahasiswa.namaDosen)
	fmt.Printf("Asal Prodi Dosen Pembimbing : %v\n", haikal.dosenPembimbingMahasiswa.prodiDosen)
	fmt.Printf("\n")

	// Deklarasi Dengan Langsung Asign Value
	dosenPembimbingAriza := dosenPembimbing{
		namaDosen:  "Pak Nuri",
		prodiDosen: "Informatika",
	}
	ariza := mahasiswa{
		namaMahasiswa:            "Ariza Akmal",
		npmMahasiswa:             "21.11.xxxx",
		dosenPembimbingMahasiswa: dosenPembimbingAriza,
	}

	// Pemanggilan Data Dengan Fungsi
	printStruct(ariza)
}
