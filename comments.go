/*
 Tanda "//" Untuk Komentar 1 Line (Bisa di samping code atau di baris baru)
 Tanda "/ dan *" Untuk Komentar lebih dari satu baris
*/

// Setting Package Golang (Untuk Default Yaitu Main)
package main

// Import Print Console atau di golang bisa pakai fmt [Bisa Pakai () ketika package lebih dari satu]
import "fmt"

// Fungsi Main yang akan di eksekusi (Pembuka Kurung {} Tidak Boleh Di awal line)
func main() {
	/*
		Setiap Fungsi yang di import dari fmt diawali dengan Huruf Besar Karena Cross Package
	*/
	fmt.Println("Hello World!")
}
