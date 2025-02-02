/*
	PACKAGE INIT Adalah sebuah konsep package / library yang menggunakan fitur init / inisialisasi dalam package nya. dalam teknik ini memungkinkan sebuah package dapat menjalankan sebuah setup function untuk melakukan persiapan defaul value, Konfigurasi System dan lain sebagainya. Dalam Golan untuk dapat menggunakan init dapat mendeklarasikan seperti ini [func init() { ini_syantax }]
*/

/*
	BLACK IDENTIFIER merupakan sebuah teknik bagaimana cara import sebuah package tanpa menggunakan function dalam fungsi lainnya. Kenapa Harus dengan Blank Identified? Karena golang ketika suatu dalam variabel atau package yang tidak digunakan dalam sebuah file go maka akan menyebabkan error. Karena golang harus mengefisiensikan sebuah memory dengan mewajibkan package / variabel digunakan sehingga tidak membuang-buang memory
*/

package main

import (
	"fmt"
	// Dapat Menggunakan Blank Identifier '_' didepan nama package agar ketika tidak diguanakan tidak akan membuat error!
	_ "github.com/HaikalRFadhilahh/learn-go-lang/package/inisialiasi"
)

func main() {
	fmt.Println("")
}
