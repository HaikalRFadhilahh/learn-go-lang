/*
	DATABASE MYSQL ini merupakan sebuah implementasi dan penggunaan dari pkg database mysql dari file package/database/db.go
*/

package main

/*
	Melakukan Import Package Mysql karena Library database/sql membutuhkan driver mysql
*/

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/HaikalRFadhilahh/learn-go-lang/package/database"
	_ "github.com/go-sql-driver/mysql"
)

/*
	Membuat Struct Untuk Handler Data Single
*/

/*
	! Ketika Data Null Pastikan Jangan Menggunakan Tipe Data Klasik. Gunakan Null Handler dari database/sql
*/

type dataMahasiswa struct {
	id              sql.NullInt64
	nama            sql.NullString
	npm             sql.NullString
	asalUniversitas sql.NullString
}

func main() {
	// Melakukan Pemanggilan Fungsi Factory Untuk mendapatkan Instance
	db, err := database.NewMysqlConnection()

	// Check Error Factory Function Database
	if err != nil {
		log.Fatalln("Database Connection Error :", err.Error())
	}

	// Membuat Fungsi Untuk Mematikan Koneksi agar tidak Memory Leak
	defer database.CloseMysqlConnection(db)

	// Deklarasi Query
	query := "SELECT nama,npm,asalUniversitas FROM mahasiswa"

	// Melakukan Query (Menggunakan method query dari database/sql) dan Tidak menggunakan param
	res, err := db.Query(query)
	if err != nil {
		log.Fatalln("Query Error :", err.Error())
	}

	// Melakukan Pembacaan Data dengan For dang fungsi sql.Rows.Next()
	fmt.Println("Data Mahasiswa !")
	for res.Next() {
		/* Membuat Variabel Untuk Menyimpan Data Sementara (Kalau Ingin Di Simpan Bisa di appent ke array Struct luar )*/
		var (
			nama            string
			npm             string
			asalUniversitas sql.NullString
		)

		/* Menggunakan sql.Rows.Scan() Untuk menerima data menggunakan Pointer*/
		/*
			! Pastikan Untuk Pass Value Scan Sesuai Dengan Data Yang DI Show Query
		*/
		if err := res.Scan(&nama, &npm, &asalUniversitas); err != nil {
			fmt.Println("Gagal Membaca Data, Error :", err.Error())
		}

		/* Show Data */
		fmt.Printf("Nama : %v, NPM : %v, Asal Universitas : %v\n", nama, npm, asalUniversitas.String)
	}

	// Break
	fmt.Println()

	// Melakukan Single Result Query menggunakan fungsi *sql.DB.QueryRow()
	// Mendeklarasikan Sebuah Struct Untuk Handler
	var haikal dataMahasiswa

	// Membuat Single Query
	singleQuery := "SELECT * FROM mahasiswa WHERE npm=? LIMIT 1"

	// Melakukan Query
	result := db.QueryRow(singleQuery, "21.11.3910")

	// Melakukan Scan Data
	if err := result.Scan(&haikal.id, &haikal.nama, &haikal.npm, &haikal.asalUniversitas); err != nil {
		fmt.Println("ada Kesalahan Saat Single Data Scan, Error :", err.Error())
	}

	// Menampilkan Data Dengan fmt
	fmt.Println("Data Mahasiswa With Single Scan Queries")
	fmt.Println("Id Data Mahasiswa : ", haikal.id.Int64)
	fmt.Println("Nama Mahasiswa : ", haikal.nama.String)
	fmt.Println("NPM Mahasiswa :", haikal.npm.String)
	fmt.Println("Asal Universitas", haikal.asalUniversitas.String)
}
