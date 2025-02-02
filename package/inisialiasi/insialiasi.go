/*
	INITIALISATION PACKAGE merupakan sebuah teknik untuk melakukan inisiasi syntax ketika sebuah file / package di import ke file lain. Untuk dapat menjalankan kumpulan syntax yang akan menjadi default exec saat di import bisa di taruh pada function bernama 'init'. Walaupun fungsi pada file ini tidak di pakai tetapi hanya di import saya fungsi init tetap terpanggil
*/

package inisialiasi

import "fmt"

/*
	Membuat Struct Untuk Simulasi Bahwa Dengan Function Init dapat melakukan Asign Default Struct
*/

type connection struct {
	typeConnetion string
	targetIP      string
}

/*
	Function Untuk Show Data Connection
*/

func (c *connection) showConnection() {
	fmt.Printf("Connection Trying To Open..\n")
	fmt.Printf("Connection Target is %v\n", c.targetIP)
	fmt.Printf("With %v Type Connection\n", c.typeConnetion)
}

/*
FUNCTION init yang akan melakukan eksekusi paling pertama ketika function di import pada file *.go yang lain
*/

func init() {
	// Declarasi
	mongoDB := &connection{}

	// Isi Variabel
	mongoDB.targetIP = "10.11.12.1"
	mongoDB.typeConnetion = "TCP"

	// Show Data Connection
	mongoDB.showConnection()
}
