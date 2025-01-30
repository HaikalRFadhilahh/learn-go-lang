/*
TYPE-CASTING dalam Go-Lang Merupakan Sebuah teknik konversi dari suatu Tipe Data Ke Tipe Data Lain. Bisa Dari Tipe Data Number Ke Tipe Data Number Lainnya, Atau Bahkan Dari String ke Int Ataupun Sebaliknya. Untuk Jenis Dan Caranya Dapat Di Jelaskan Di Bawah Ini :
1. Int & Float => bisa dengan cara nama_variabel_int := int(variabel_float) / nama_variabel_float := float32(int) / := float64(int) PERLU DI PERHATIKAN BAHWA KONVERSI DARI FLOAT KE INT DI GENAPKAN KE BAWAH!.
2. String & Number (int & Float32) => Bisa dengan menggunakan library 'strconv' dengan fungsi Atoi(nama_variabel_string) untuk konversi dari string -> int . Untuk Int -> String Bisa menggunakan Itoa(nama_variabel_int) . Dan Untuk ParseFloat(nama_variabel_float) String -> float dan FormatFloat(nama_variabel) untuk Float -> string
*/
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
FUNCTION Untuk Melakukan Konversi Dari INT Ke Float32 / Float64
*/
func convertIntToFloat32(number int) {
	fmt.Printf("Data Dari '%v' Value '%v' Menjadi '%v' Value '%v'\n", reflect.TypeOf(number), number, reflect.TypeOf(float32(number)), float32(number))
}

func convertIntToFloat64(number int) {
	fmt.Printf("Data Dari '%v' Value '%v' Menjadi '%v' Value '%v'\n", reflect.TypeOf(number), number, reflect.TypeOf(float64(number)), float64(number))
}

/*
FUNCTION untuk Melakukan Konversi Dari Float32 / 64 ke INT
*/
func convertFloatToInt(number float64) {
	fmt.Printf("Data Dari '%v' Value '%v' Menjadi '%v' Value '%v'\n", reflect.TypeOf(number), number, reflect.TypeOf(int(number)), int(number))
}

/*
DAFTAR FUNCTION Untuk Convert String ke Number (Int / Float32 / Float64) Dengan Library 'strconv'
*/
func convertStringToInt(str string) {
	v, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Ada Kesalahan Saat Konversi Nilai")
	} else {
		fmt.Printf("Data Dari '%v' Value '%v' Menjadi '%v' Value '%v'\n", reflect.TypeOf(str), str, reflect.TypeOf(v), v)
	}
}

func convertStringToFloat(str string) {
	v, err := strconv.ParseFloat(str, 32)
	if err != nil {
		fmt.Println("Ada Kesalahan Saat Konversi Nilai")
	} else {
		fmt.Printf("Data Dari '%v' Value '%v' Menjadi '%v' Value '%v'\n", reflect.TypeOf(str), str, reflect.TypeOf(v), v)
	}
}

func convertIntToString(number int) {
	v := strconv.Itoa(number)
	fmt.Printf("Data Dari '%v' Value '%v' Menjadi '%v' Value '%v'\n", reflect.TypeOf(number), number, reflect.TypeOf(v), v)
}

func convertFloatToString(number float64) {
	v := strconv.FormatFloat(number, 'f', 2, 32)
	fmt.Printf("Data Dari '%v' Value '%v' Menjadi '%v' Value '%v'\n", reflect.TypeOf(number), number, reflect.TypeOf(v), v)
}

func main() {
	// Berikan Nilai Untuk testing
	nilaiInt := 30
	nilaiFloat32 := 30.20
	nilaiIntString := "70"
	nilaiFloatString := "30.90"

	// Convert Int To Float32
	convertIntToFloat32(nilaiInt)

	// Convert Int To Float64
	convertIntToFloat64(nilaiInt)

	// Convert Float64 Ke Int
	convertFloatToInt(float64(nilaiFloat32))

	// Convert String To Int
	convertStringToInt(nilaiIntString)

	// Convert String To Float
	convertStringToFloat(nilaiFloatString)

	// Convert Int To String
	convertIntToString(nilaiInt)

	// Convert Float64 To String
	convertFloatToString(float64(nilaiFloat32))
}
