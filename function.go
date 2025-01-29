/*
FUNCTIONS merupakan teknik untuk memecah kode menjadi bagian yang lebih kecil sehingga kode tidak berkumpul dalam fungsi main / fungsi utama sehingga mudah di kembangkan dan di baca oleh orang lain ketika ikut mengembangkan project. Ada Beberapa tipe fungsi dalam go-lang yang dapat di manfaatkan untuk menata dan mengefisiensikan kode menjadi lebih baik :
1. Call Function / Function No Paramater & No Return adalah fungsi untuk memanggil sesuatu namun tanpa paramater atau nilai yang diterima tetapi juga tidak mengembalikan nilai sehingga saat melakukan eksekusi fungsi tidap perlu menampung hasil dari return function
2. Paramater Function No Return merupakan suatu fungsi yang tidak memiliki return nilai seperti fungsi sebelumnya namun menerima paramater yang nilainya akan digunakan didalam fungsi tersebut
3. Return Function merupakan fungsi yang mengembalikan nilai dan bisa di olah lagi difungsi yang lain, tipe fungsi ini juga bisa menerima paramater atau tidak namun wajib mengembalikan nilai sesuai yang deklarasikan return tipe data fungsinya.
4. Recursion Function merupakan suatu teknik pembuatan fungsi yang dapat memanggil dirinya sendiri sehingga lebih efisien namun sangat jarang kasus yang menggunakan ini
! Harap berhati-hati saat menggunakn Recursion Function karena ketika tidak ada exit condition akan memberatkan memory / RAM
*/

package main

import "fmt"

/*
Call Function / Function No Return No Paramater yang akan dicontohkan hanya menampilkan Console Output Sehingga Tidak Return Value Dan Tidak menerima Paramater
*/
func sayHello() {
	fmt.Println("Hello Everyone, Haikal In There and Learn Go-Lang")
}

/*
Paramater Function No Return adalah fungsi yang dapat menerima paramater namun tidak mengembalikan value yang dicontohkan menerima paramater string nama dan nama universitas yang akan di gunakan pada blok fungsi
*/
func introduction(namaMahasiswa string, namaUniversitas string) {
	fmt.Printf("Nama Saya %v Dari Universitas %v\n", namaMahasiswa, namaUniversitas)
}

/*
Return Function adalah Fungsi yang mengembalikan Value namun dapat menerima paramater dan tidak dicontohkan yaitu fungsi untuk mendeteksi apakah nilai paramater yang dikirimkan angka genap atau bukan
*/
func isEvenNumber(angka int) bool {
	if (angka % 2) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	// Call Function
	sayHello()

	// Function No Return
	introduction("Haikal R Fadhilah", "AMIKOM Yogyakarta")

	// Function Return Value
	isEven := isEvenNumber(5)
	if isEven {
		fmt.Println("Ini Adalah Angka Genap")
	} else {
		fmt.Println("Ini Adalah Angka Ganjil")
	}
}
