/*
	INTERMEDIATE SECTION
	POINTER dalam Go-Lang merupakan sebuah teknik untuk melakukan perubahan data secara langsung tanpa melakukan replikasi atau lebih dikenal dengan istilah 'pass by value' dengan membuat sebuah variabel pointer sebuah variabel akan menyimpan sebuah alamat dari value / variabel yang di pasang, Deklarasi Variabel Pointer dapat di Declare Secara Inference atau Manual menggunakan var tetapi pastikan menggunakan tanda asterik / * diawal tipe_data nya sehingga memberikan tanda bahwa variabel tersebut menyimpan alamat memori bukan value.
*/

/*
	Dalam implementasi Pointer terdapat 2 tanda yang dapat digunakan untuk dapat mengoptimal kan fungsi pointer, Berikut tanda operator yang dapat digunakan :
	1. Asterik Operator / '*' berguna untuk mendeklarasikan variabel bahwa itu pointer dengan cara deklarasi '*tipe_data' dan bermanfaat untuk mengambil value ketika sebuah tipe data bertipe pointer, karena jika sebuah variabel pointer diprint tanpa asterik maka akan mencetak lokasi memori bit bukan value aslinya.
	2. Reference Operator / '&' berfungsi untuk mengekstrak alamat memori value dari variabel atau secara langsung. dengan syarat variabel masih berupa non pointer. Tanda Reference tidak akan berfungsi jika variabel sudah bertipa * / pointer
*/

/*
	! Variabel Pointer Ketika Deklarasikan Secara Eksplisit dan Tidak Ada Nilai Default maka akan Bernilai Nill / Kosong
	! Jika Membuat Pointer Dari Struct namun Atribute nya bukan Pointer maka untuk akses value dari pointer sruct tidak usah pakai asterik. Namun ketika Struct dan Atribute Nya Pointer maka wajib menggunakan asterik.
*/

package main

import "fmt"

/*
	Membuat Struct Untuk Simulasi Pointer
*/

type sekolahTinggi struct {
	nama              string
	tipeSekolahTinggi string
	lokasi            *string
}

/*
	FUNCTION Main
*/

func main() {
	// Contoh Simulasi Pass By Value / Non Pointer
	var namaMahasiswa string = "Haikal Raditya Fadhilah"
	var copyNamaMahasiwa string = namaMahasiswa

	fmt.Println("Simulasi Pass By Value / Non Pointer")
	fmt.Printf("Isi Variabel lama namaMahasiswa Yaitu : %v\n", namaMahasiswa)
	fmt.Printf("Isi Variabel lama copyNamaMahasiswa Adalah : %v\n", copyNamaMahasiwa)

	copyNamaMahasiwa = "Yoko"

	fmt.Println("Isi Variabel namaMahasiswa Yang Telah Dirubah dari copyNamaMahasiswa adalah :", namaMahasiswa)
	fmt.Println("Isi Variabel copyMahasiswa Yang Telah Dirubah : ", copyNamaMahasiwa)
	/*
		Dari Hasil Yang ditampilkan bahwa variabel namaMahasiswa tidak berubah walaupun copyNamaMahasiswa berubah di karenakan variabel copyNamaMahasiswa melakukan copy value bukan merujuk pada pointer yang sama sehingga ketika variabel copyNamaMahasiswa di rubah tidak akan mempengaruhi namaMahasiswa dan variabel tersebut menjadi aman dengan Pass By Value
	*/

	fmt.Println()

	// Contoh Simulasi Pass By Referece / Pointer
	var namaUniversitas string = "Universitas Amikom Yogyakarta" // Deklarasi Variabel Secara Langsung
	var ptrNamaUnversitas *string = &namaUniversitas             // Menggunakan & Untuk Ekstrak Lokasi Alamat Memory value variabel namaUniversitas

	fmt.Printf("Isi Variabel namaUniversitas : %v\n", namaUniversitas)
	fmt.Printf("Isi Variabel ptrNamaUniversitas (pointer var namaUniversitas) : %v \n", *ptrNamaUnversitas) // Menggunakan Asterik Karena Untuk mengambil Value karena variabel ptrNamaUniverstas adalah Pointer (Ingat Keterangan Di Atas)

	// Merubah Isi Variabel ptrNamaUniversitas
	*ptrNamaUnversitas = "Universitas Gadjah Mada" // Menggunakan Asterik Karena Kita akan Merubah Valuenya Bukan Alamatnya

	// Menampilkan Data Yang Telah Di Rubah
	fmt.Println("Isi Variabel namaUniversitas Yang Telah Dirubah dari ptrNamaMahasiswa : ", namaUniversitas)
	fmt.Println("Isi Variabel ptrNamaUniversitas Yang Telah Dirubah : ", *ptrNamaUnversitas)
	fmt.Println()

	// Implementasi Struct Pointer
	lokasiAmikom := "Jalan RingRoad Utata"
	amikom := &sekolahTinggi{
		nama:              "Universitas Amikom Yogyakarta",
		tipeSekolahTinggi: "Universitas",
		lokasi:            &lokasiAmikom,
	}

	fmt.Println("Nama Sekolah Tinggi : ", amikom.nama)              // Walapun Struct Amikom Pointer namun akses nama tidak perlu pakai * karena atribute nama tidak pointer
	fmt.Println("Jenis Sekolah Tinggi :", amikom.tipeSekolahTinggi) // Walapun Struct Amikom Pointer namun akses tipeSekolahTinggi tidak perlu pakai * karena atribute tipeSekolahTinggi tidak pointer
	fmt.Println("Lokasi Sekolah Tinggi : ", *amikom.lokasi)         // Akses lokasi perlu asterik karena atribute lokasi adalah pointer
}
