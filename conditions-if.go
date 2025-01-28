/*
	CONDITIONS IF ELSE merupakan kondisi untuk mengeksekusi blok kode tertentu. Sebuah blok kode akan eksekusi ketika nilai kondisi bernilai true atau benar. IF-Else dapat memiliki banyak kondisi dan bisa lebih dari satu. Ada beberapa tipa IF-Else yang tersedia dalam sebua bahasa pemrograman yaitu :
	1. IF merupakan satu kondisi tertentu ketika kondisi yang di wajibkan terpenuhi atau true maka blok kode akan di eksekusi
	2. IF - Else Merupakan kondisi ketika terpenuhi akan eksekusi block kode namun ketika tidak terpenuhi akan ada block kode lain yang akan menghandle exception
	3. IF- Else If - Else Merupakan Kode Kondisi yang memiliki berbagai kondisi untuk handle berbagai kebutuhan kode dan requirements
	4. Nested IF (IF-Else Bersarang) Merupakan kode yang terdapat IF-Else structure didalam IF Else
*/

package main

import "fmt"

func main() {
	// Credentials
	username := "admin"
	password := "password"

	// IF Condition
	if username == "admin" {
		fmt.Println("Username Sesuai, Dan Kode IF Kondisi Berhasil Di Eksekusi")
	}

	// IF Else Condition
	username = "haikal"
	if username == "admin" {
		fmt.Println("Kode Block IF Else berhasil di eksekusi, Karena Username Admin")
	} else {
		fmt.Println("Kode Block IF Else Berhasil Di Eksekusi, namun masuk ke block Else Karena tidak memenuhi kondisi IF")
	}

	// IF - Else If Condition
	username = "admin"
	if username == "admin" && password == "password" {
		fmt.Println("Kode Block IF - Else IF codition true berhasil di eksekusi. Username Dan Password Sesuai")
	} else if username == "admin" && password != "password" {
		fmt.Println("Kode IF Else IF Berhasil masuk ke Else If Condition, Karena Username Benar Namun Password Salah!")
	} else {
		fmt.Println("Kode IF Else IF berhasil di eksekusi namun masuk kode Else, Username Dan Password tidak ada yang sesuai!")
	}

	// Nested IF Condition
	if username == "admin" {
		if password == "password" {
			fmt.Println("Username Dan Password Benar, Kode Nested IF Berhasil dan Sukses Di Eksekusi")
		} else {
			fmt.Println("Kode Nested Else Berhasil di eksekusi, Namun Username Benar tetapi Password Salah!")
		}
	}
}
