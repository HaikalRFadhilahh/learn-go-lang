/*
	STRUCT-METHOD-POINTER merupakan sebuah teknik untuk mengubah value menggunakan struct method dari function, yang memungkinkan ketika sebuah struct method di panggil dan mengganti value dari atrribute maka Variabel Struct yang ada diluar function akan berubah, Secara konsep mirip dengan Function namun saay melakukan paramater reciever pada method kita menggunakan sebuah pointer ataupun tanda asterik.
*/

package main

import "fmt"

/*
MEMBUAT STRUKTUR STRUCT
*/
type laptop struct {
	merek        string
	warna        string
	supportTypeC bool
}

/*
	Membuat Struct Method Dengan Pointer Reciever
*/

func (l *laptop) setMerekLaptop(m string) {
	l.merek = m
}

func (l *laptop) setWarnaLaptop(w string) {
	l.warna = w
}

func (l *laptop) setSupportTypeC(s bool) {
	l.supportTypeC = s
}

/*
	Membuat Struct Method Tanpa Pointer Variabel Reciever
*/

func (l laptop) setMerekLaptopPBV(m string) {
	l.merek = m
}

func (l laptop) setWarnaLaptop2PBV(w string) {
	l.warna = w
}

func (l laptop) setSupportTypeCPBV(s bool) {
	l.supportTypeC = s
}

/*
	Membuat Function Untuk Menampilkan Seluruh Value Dari Atribute Struct
*/

func showDataLaptop(l laptop) {
	fmt.Println("Merek Laptop Adalah : ", l.merek)
	fmt.Println("Warna Laptop Adalah : ", l.warna)
	fmt.Println("Apakah Laptop Support Type-C : ", l.supportTypeC)
}

func main() {
	// Deklarasi Struct
	macbookAirM1 := laptop{}
	asusTufGaming := laptop{}

	// Show Default Value
	showDataLaptop(macbookAirM1)
	showDataLaptop(asusTufGaming)

	// Exec Struct Method
	// [Pointer Struct Function]
	macbookAirM1.setMerekLaptop("Apple Macbook Air M1 2020")
	macbookAirM1.setWarnaLaptop("Silver")
	macbookAirM1.setSupportTypeC(true)

	// [Non Pointer Struct Method]
	asusTufGaming.setMerekLaptopPBV("Asus TUF Gaming")
	asusTufGaming.setWarnaLaptop2PBV("Black")
	asusTufGaming.setSupportTypeCPBV(false)

	// Show New Value After Using Pointer
	fmt.Println("Data Value Struct Macbook Air M1 2020")
	showDataLaptop(macbookAirM1)
	fmt.Println()
	fmt.Println("Data Value Struct Asus TUF Gaming")
	showDataLaptop(asusTufGaming)

}
