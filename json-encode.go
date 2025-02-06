/*
JSON atau Javascript Object Notation seperti yang telah saya sebutkan sebelumnya pada materi decode. Bahwa untuk melakukan sebuah konversi json entah Encode / decode kita bisa menggunakan []byte den struct. Kebanyakan Function Pada library bawaan golan json mewajibkan untuk menggunakan Pass By Pointer dan juga []byte dan kita konversi sendiri menjadi Struct ataupun Map dengan cara Type Assertion / Conver Tipe Data.
! Cara Konveri Sebuah Struct Dengan library Json ada 2 cara yaitu marshal dan NewEncoder().Encode()
! Untuk lebih jelas terkait dengan Paramater dan Return yang dihasilkan Bisa Lihat Di Dokumentasi go
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

/*
	Membuat Struct Untuk simulasi Encoding Json Dari Struct
*/

type orang struct {
	Nama        string `json:"nama"`
	Umur        int    `json:"umur"`
	TempatLahir string `json:"tempatLahir"`
}

/*
FUNCTION MAIN HANDLER
*/

func main() {
	// Memmbuat Sebuah Variabel Struct
	haik := orang{"Haikal Raditya Fadhilah", 21, "Bantul"}

	// Melakukan Konversi Dengan Json Marshal (Return 2 Paramater yaitu []byte dan error untuk check kondisi)
	hasil, err := json.Marshal(&haik)
	if err != nil {
		log.Fatalln("Decoding Has Error, The Errors Is :", err.Error())
	}

	// Melihat Hasil Encoding Struct
	fmt.Println("Konversi Json Dari Struct Dengan Marshal!")
	fmt.Println("Hasil Encoding Struct Ke Array of Byte ([]byte) :", hasil)
	fmt.Println("Hasil Encoding Struct ke String :", string(hasil))

	// Break
	fmt.Println()

	// Melihat Hasil Encoding Struct
	fmt.Println("Hasil Konversi Struct Ke Json dengan NewEncoder os.Stdout")

	// Melakukan Encode Json Dengan NewEncoder // Langsung Output Karena Langsung Output Console atau Write Dari io.Writer
	// ! Kalau Mau Di Encode Pakai Marshal Aja
	if err := json.NewEncoder(os.Stdout).Encode(&haik); err != nil {
		fmt.Println("Encoding Struct To Json Error, The Error :", err.Error())
	}

}
