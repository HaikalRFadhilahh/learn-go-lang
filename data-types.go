/*
	Tipe Data (Data Type) Tipe data merupakan pembeda tipe data yang akan di simpan ke variabel ada 2 jenis yaitu tipe data klasik dan kompleks.
	Tipe Data:
	- int / int8 / int32 / int64 : untuk menyimpan angka positif maupun negatif
	- float32 / float64 : Untuk Menyimpan Angka Desimal atau koma mau itu positif maupun negatif
	- uint / uint8 / uint32 / uint64 : Untuk Menyimpan Angka Positif
	- bool : Untuk menyimpan Boolean (True / False)
	- string : Untuk menyimpan kata atau kalimat
	Alias :
	- rune : Alias dari int32 biasanya untuk menyimpan char / single character, exp :'A'
	- byte : Alias dari uint8
*/

package main

import "fmt"

func main() {
	// int
	var intDefault int
	intValue := 20

	// float32
	var float32Default float32
	float32Value := 3.86

	// uint
	var uintDefault uint
	uintValue := 5

	// bool
	var boolDefault bool
	boolValue := true

	// string
	var stringDefault string
	stringValue := "Haikal"

	// rune
	var runeDefault rune
	runeValue := 'P'

	// byte
	var byteDefault byte
	byteValue := 3

	// Print Data Types
	fmt.Println("Int Default :", intDefault)
	fmt.Println("Int Value :", intValue)

	fmt.Println("float32 Default :", float32Default)
	fmt.Println("float32 Value :", float32Value)

	fmt.Println("uint Default :", uintDefault)
	fmt.Println("uint Value :", uintValue)

	fmt.Println("bool Default :", boolDefault)
	fmt.Println("bool Value :", boolValue)

	fmt.Println("string Default :", stringDefault)
	fmt.Println("string Value :", stringValue)

	fmt.Println("rune Default :", runeDefault)
	fmt.Println("rune Value :", runeValue)

	fmt.Println("byte Default :", byteDefault)
	fmt.Println("byte Value :", byteValue)
}
