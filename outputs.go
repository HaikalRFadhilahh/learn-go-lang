/*
	Outputs adalah fungsi dari package fmt default Golang untuk print ke terminal active. Beberapa fungsi bisa di manfaatkan untuk logging bahkan trace error. Berikut fungsi dan kegunaan dari fungsi fmt :
	1. Print() : merupakan perintah untuk menampilkan text atau value variabel tanpa newline di akhir atau '\n', value yang di kiirm bisa lebih dari satu di separate by ,
	exp : fmt.Print("Haikal","Raditya","Fadhila","dll")
	2. Println() : Memiliki fungsi yang sama seperti Print() namun akan menambahkan newline di akhir
	exp : fmt.Println("Saya","Baik") // nanti setelah kata "Saya Baik" akan ada newline yang terbentuk dari fungsi Println()
	3. Printf() : memiliki fungsi sama seperti fungsi Print sebelumnya namun dapat di format dengan variabel atau value lain dengan menambahkan %v untuk posisi yang ingin di replace
	exp : fmt.Printf("%v Berkuliah Di %v","Haikal","UGM")
	output : Haikal Berkuliah Di UGM
	! Dalam Fungsi Printf() dapat di format untuk extract informasi tertentu, dapat di lihat pada url : https://www.w3schools.com/go/go_output.php
*/

package main

import "fmt"

func main() {
	// Print()
	fmt.Print("Saya")
	fmt.Print("Haikal Fadhilah")
	// Output : SayaHaikalFadhilah

	// Menambahkan Space Dengan Print()
	fmt.Print("\n")

	// Println()
	fmt.Println("Saya Kuliah di UGM")

	// Printf()
	fmt.Printf("%v Tinggal Di %v", "Haikal", "Yogyakarta")
}
