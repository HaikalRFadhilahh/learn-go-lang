/*
	DATABASE Merupakan sebuah tempat penyimpanan sebuah data agar data yang di masukkan user tidak hilang. Namun untuk proses memasukkan data kedalam database di perlukan sebuah connection / teknik dari sebuah aplikasi kita. Dalam kasus kita akan menggunakan MySql atau Relation Database. Dalam menggunakan mysql kita harus melakukan sebuah init connection dan object untuk melakukan aksi CRUD (Create, Read, Update, Delete).

	Sebelum melakukan simulasi koneksi ke database mysql di perlukan library yang perlu di install menggunakan `go get` dengan perintah di bawah ini :
	! go get -u github.com/go-sql-driver/mysql
	! Walaupun untuk membuat sebuah koneksi menggunakan package golang bawaan yaitu `database/sql` kita tetap perlu mengimport dengan blank identifier pada package main / yang menggunakan instance

	Dalam Membuat Koneksi menggunakan sql.Open() membutuhkan param tipe_database,Connection String dengan format sebagai berikut :
	"username:password@tcp(host:port)/nama_database"
*/

package database

import (
	"database/sql"
	"fmt"
)

/*
	Membuat Struct Untuk Data Koneksi Mysql
*/

type databaseConnection struct {
	username     string
	password     string
	host         string
	port         string
	databaseName string
}

var dbConn databaseConnection

/*
Membuat Sebuah Koneksi (Sebaiknya Credentials Di Masukkan Kedalam .env dan di import)
*/
func init() {
	dbConn.host = "127.0.0.1"
	dbConn.port = "3306"
	dbConn.username = "root"
	dbConn.password = "1235813"
	dbConn.databaseName = "learn-golang"
}

/*
	Membuat Contructor Function / Factory Function Untuk Membuat Instance Database Mysql
*/

func NewMysqlConnection() (*sql.DB, error) {
	// Membuat Connection String
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbConn.username, dbConn.password, dbConn.host, dbConn.port, dbConn.databaseName)

	// Membuat Object Database Connection menggunakan database/sql
	/*
		! Err pada kasus ini merupakan error ketika melakukan initiliasi string, bukan error koneksi
	*/
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return db, err
	}

	/*
		! Fungsi Ping Untuk Melakukan Test Koneksi jika Error Langsung Return Interface Error, Jika Tidak Return nil
	*/

	return db, db.Ping()
}

/*
	Membuat Fungsi Untuk Mematikan Koneksi
*/

func CloseMysqlConnection(db *sql.DB) {
	db.Close()
}
