/*
	ORM atau Object-Relation-Mapping merupakan sebuah teknik untuk mempermudah programmer dalam melakukan pengembang terhadap database. ORM dalam golang untuk mysql yaitu GORM salah satu dan satu-satunya saat ini. Dengan menggunakan GORM kita dapat lebih mudah untuk melakukan pengambilan data dan tanpa harus membuat query secara manual.
*/

/*
	Sebelum melakukan pengembangan menggunakan gorm kita harus melakukan instalasi terlebih dahulu menggunakan perintah go get dari golang. Berikut merupakan script untuk menginstall library gorm dan driver mysql:
	`go get -u gorm.io/gorm`
	`go get -u gorm.io/driver/mysql`
*/

/*
	Dalam menggunakan gorm kita memiliki banyak opsi tambahan seperti dapat melakukan sebuah migrasi dan dapat menggunakan tag untuk hal tersebut.
*/

/*
	Untuk meningkatkan Fungsi Dari Gorm Kita Juga Dapat menggunakan Json Tag Models Dari Gorm Agar Dapat Melakukan Konfigurasi Dengan Lebih AdvanceN
*/

package orm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
	Membuat Variabel Untuk Credentials Koneksi
*/

var (
	host     string
	port     string
	username string
	password string
	dbName   string
)

/*
	Membuat Init Function Untuk Koneksi Gorm
*/

func init() {
	host = "127.0.0.1"
	port = "3306"
	username = "root"
	password = "1235813"
	dbName = "learn-golang"
}

/*
	Mendeklarasikan Fungsi Init / New Untuk Membuat Object Dari Gorm.DB
*/

func NewGormConnection() (*gorm.DB, error) {
	// Membuat Database Connection String
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
