/*
	ERROR dalam golang merupakan sebuah interface untuk melakukan handle error dan dapat memberikan sebuah tanda / informasi bahwa sebuah fungsi yang akan di eksekusi apakah berjalan dengan baim atau tidak. Dengan memanfaatkan multiple Return Golang yaitu Sebuah Value dan Error maka kita bisa melakukan handling error dengan baik.
	! error dan errors itu beda. Jika error itu adalah interface dari golang untuk implementasi error dan errors merupakan sebuah package yang dapat membuat error, membandingkan error, dll
	Interface Error Yaitu Sebagai Berikut :
	`type error interface {
		Error() string
	}`
	! Sebuah interface error memiliki Fungsi Error() yang return nilai string errornya
*/

/*
	Untuk Implementasi Nyata error kita dapat menggunakan berbagai Tipe Error sebagai Handling Yang Berbeda bisa menggunakan Type Assertion
*/

/*
	Untuk Membuat Error Kita Bisa Menggunakan 2 Cara Yaitu Custom Error Dan Package errors.New.
	Pada file ini kita akan menggunakan generate error secara otomatis dulu dengan package errors
*/

package main

import (
	"errors"
	"fmt"
)

/*
	Membuat Function Dengan Multiple Return Apakah Error Atau Tidak
*/

func simulasiError(s bool) error {
	if s {
		return nil
	}

	// Dengan Menggunakan errors.New error return akan membalikkan struct yang otomatis dibuatkan oleh package error
	return errors.New("Internal Server Error")
}

func main() {
	if err := simulasiError(false); err != nil {
		// Fungsi Error() pasti ada dalam struct error karena dari golang mewajibkan untuk interface fungsi Error() denganr return string
		fmt.Println("Terdapat Error :", err.Error())
	}
}
