/*
	ANONYMOUS FUNCTION merupakan sebuah teknik untuk mendeklarasikan sebuah fungsi namun tanpa pemberian nama pada fungsi. Dengan Anonymous Function Developer memungkinkan untuk dapat melakukan eksekusi secara langsung saat itu, bahkan dapat mengirimkan paramater untuk fungsi tersebut. Layaknya javascript kita juga dapat Memasukkan Anonymous Function Ke dalam variabel Sepert Deklarasi Function JS dengan ES6. Berikut Beberapa Cara Deklarasi Anonymous Function :
	1. Anonymous Function Langsung Eksekusi [func (param1 tipe_data,...){}(param1,...)] Dengan Cara ini anonymous function akan langsung di eksekusi pada saat itu juga
	2. Anonymous Function Dimasukkan Ke Variabel [nama_variabel := func(param1 tipe_data,...){}] Dengan Cara ini anonymous function di masukkan ke dalam variabel dan ketika ingin di eksekusi dengan cara [nama_variable(param1,...)]
*/

package main

import "fmt"

/*
DEKLARASI ASIGN ANONYMOUS FUNCTION KE VARIABEL
*/
var funcSayHi = func() {
	fmt.Println("Hi Users, HEHEHE..")
}

func main() {
	// Anonymous Function Non Asign Variabel
	func(a int, b int) {
		fmt.Printf("Hasil Perkalian Angka '%v' * '%v' Adalah : %v\n", a, b, a*b)
	}(10, 5)

	// Execute Anonymouse Function di Variabel
	funcSayHi()
}
