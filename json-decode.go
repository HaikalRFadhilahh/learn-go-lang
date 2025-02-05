/*
JSON atau JAVASCRIPT OBJECT NOTATION merupakan sebuah teknik pengiriman data antar aplikasi atau biasa diimplementasikan dengan konsep REST API (Resftul Application Programming Interface) dalam golang sebuah Json dapat di create atau di Decode serta encode menggunakan package bawaan dari golang yaitu json. Sebuah Json dapat decode maupun di encode menggunakan sebuah struct namanya json struct tag. adapun cara mendeklarasikan sebuah struct dengan memberi tag `json:"nama_json"`
*/

/*
Pada Dasarnya Untuk Melakukan Decode Sebuah json / json-string / json-[]byte dapat menggunakan fungsi json.Unmarshal([]byte,any)
*/

/*
! Json Unmarshal Hanya Dapat Di Konversi Menjadi map[string]string ketika defaultnya Interface{} / any
! Atau
! Bisa menggunakan Struct Yang Sesuai Nama Kolomnya Dan Menggunakan Json Tag jika tidak ada JSON Tag Maka diambil dengan sesuai Nama Variabel
*/
package main

/*
	Melakukan Import Library Json dari Default Golang
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

/*
	Membuat Struct Untuk Json Handler Decode
	! Atribute Harus Di Export dengan Awal Huruf Kapital untuk merubah Decode atau Encode bisa menggunakan `json: "nama_konversi_json"`
	Ada teknik spesial dalam json tag yaitu 'omitempty' agar menghilangkan sebuah atribute json ketika sebuah atribute memiliki nilai default dari tipe data
*/

type BiodataMahasiswa struct {
	Nama            string `json:"nama"`
	Umur            int    `json:"umur"`
	AsalUniversitas string `json:"asalUniversitas"`
}

/*
	Membuat Struct Untuk Simulasi Error Dengan json NewDecoder
*/

type BiodataMahasiswaErr struct {
	Nama string `json:"nama"`
	Umur int    `json:"umur"`
	//AsalUniversitas string `json:"asalUniversitas"` // Kalau Pakai Json Unmarshal / NewDecoder().Decode() tetap aman
	// ! Namun Akan Error Jika Harus di validasi bahwa semua struct bersifat required.
}

/*
	MAIN FUNCTION
*/

func main() {
	// Membuat Sebuah String Json yang telah di encode dari string
	var contohJson string
	contohJson = `{"nama": "Haikal Raditya Fadhilah","umur" : 21,"asalUniversitas": "Universitas Amikom Yogyakarta"}`

	// Hasil Decode Dari Json Unmarshal
	var resultJson any

	// Konversi Dengan Interface{} / Any yang Return Nya map[string]string
	err := json.Unmarshal([]byte(contohJson), &resultJson)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Konversi Dengan Struct yang memiliki tag json
	/* Deklarasikan Dulu Variabel Structnya */
	var dataMahasiswa BiodataMahasiswa

	/* Lakukan Konversi dengan json Unmarshal (Kelebihan Bisa di Akses Dengan Tipe Data Struct /  Akses ke Atribute) */
	if err := json.Unmarshal([]byte(contohJson), &dataMahasiswa); err != nil {
		log.Fatalln(err.Error())
	}

	// Konversi Dengan fungsi json.NewDecoder().Decode()
	/* Deklarasi Variabel Struct */
	var mahasiswaBiodata BiodataMahasiswaErr

	/* Simulasi Error */
	reader := json.NewDecoder(strings.NewReader(contohJson))
	reader.DisallowUnknownFields()
	err = reader.Decode(&mahasiswaBiodata)
	if err != nil {
		fmt.Println("Error Decode With NewEncoder and DisallowUnkownFields, Err : ", err.Error())
	} else {
		fmt.Println("Hasil Konveri Json Dengan json NewEncoder dan DisallowUnkownFields")
		fmt.Println("Nama Mahasiswa :", mahasiswaBiodata.Nama)
		fmt.Println("Umur Mahasiswa :", mahasiswaBiodata.Umur)
	}

	// Break
	fmt.Println()

	// Print hasil Result Dengan Konversi Interface{} / any
	fmt.Println("Hasil Dari Koversi Json Dengan Variabel Interface Kosong / Any")
	fmt.Println(resultJson)

	// Break
	fmt.Println()

	// Print Hasil Result Dengan Struct Pointer
	fmt.Println("Hasil Konversi Json Dengan Struct dan Json Tag")
	fmt.Println("Nama Mahasiswa :", dataMahasiswa.Nama)
	fmt.Println("Umur Mahasiswa :", dataMahasiswa.Umur)
	fmt.Println("Asal Universitas Mahasiswa :", dataMahasiswa.AsalUniversitas)
}
