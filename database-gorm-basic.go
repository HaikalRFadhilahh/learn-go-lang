/*
	GORM merupakan Sebuah ORM untuk mempermudah dalam melakukan pengembangan aplikasi yang memerlukan Transfer dan Get data dari database mysql. Pada File Ini Merupakan Implemnetasi Penggunaan Gorm Yang Di Support Local Package dari `package/orm/db.go`
*/

package main

import (
	"log"

	"github.com/HaikalRFadhilahh/learn-go-lang/package/orm"
)

func main() {
	// Melakukan Init Gorm Connection
	_, err := orm.NewGormConnection()

	// Melakukan Pengecekan Error Ketika Database Tidak Terhubung
	if err != nil {
		log.Fatalln("Database Connection Error :", err.Error())
	}

	// Melakukan Auto Migrate Ketika Database Table Tidak Ditemukan

}
