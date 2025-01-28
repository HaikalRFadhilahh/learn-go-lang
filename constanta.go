/*
	Konstanta adalah variabel yang tidak dapat di ubah ubah setelah pertama kali di asign, contoh deklarasi
	const nama_variabel = value_variabel
	! Gunakan Dengan Tepat
*/

package main

import "fmt"

func main() {
	// Deklarasi Pertama Konstanta
	const namaUniversitas = "Universitas Amikom Yogyakarta"
	// Merubah Konstanta
	// namaUniversitas = "UGM" // Akan Error

	// Menampilkan Data Variabel Konstanta
	fmt.Println("Nama Universitas :", namaUniversitas)
}
