/*
	TYPE SWITCHES dalam Go-Lang Merupaakn teknik untuk melakukan konversi tipe data menggunakan Switch case, dengan cara swicth kita bisa membuat case untuk mencocokan apakah tipe yang sesuai dari value Interface{} / any setelah case ditemukan dan tipe data cocok, Maka dapat di konversi dengan baik. Adapun case default yang akan di eksekusi ketika tipe data / case yang di konfigurasi tidak ada yang cocok sehingga default akan return nilai defaultnya yaitu any / interface{}.
	! Mungkin TYPE SWITCH Jarang Digunakan pada Proejct Nyata. Namun Gunakan Fitur ini Jika Memungkinkan Dan ketika Cara ini lebih efisien daripada Type Assertion Manual.
*/

package main

import (
	"fmt"
)

/*
Membuat Struct Untuk Melakukan Simulasi TYPE SWITCHES Pada Fungsi Main
*/
type handphone struct {
	merek string
	warna string
	harga int
}

/*
	MAIN FUNCTION
*/

func main() {
	// Buat Tipe Data Any / Interface{}
	var namaUniversitas any = "Universitas Amikom Yogyakarta"

	// Melakukan Type Switches
	switch namaUniversitas.(type) {
	case int:
		fmt.Println("Berhasil Melakukan Konversi, Tipe Data dari Variabel namaUniversitas Adalah : Int")
	case string:
		fmt.Println("Berhasil Melakukan Konversi, Tipe Data dari Variabel namaUniversitas Adalah : String")
	case bool:
		fmt.Println("Berhasil Melakukan Konversi, Tipe Data dari Variabel namaUniversitas Adalah : Boolean")
	case float32:
		fmt.Println("Berhasil Melakukan Konversi, Tipe Data dari Variabel namaUniversitas Adalah : float32")
	case float64:
		fmt.Println("Berhasil Melakukan Konversi, Tipe Data dari Variabel namaUniversitas Adalah : float64")
	default:
		fmt.Println("Tidak Dapat Mengkonversi Tipe Data Any Ke Daftar Tipe Data Yang Di Konfigurasi")
	}
}
