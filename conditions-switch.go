/*
	CONDITIONS SWITCH meruapakn teknik handle kondisi namun dengan cara yang berbeda dan kurang fleksibel. Namun sangat cocok untuk handle string dan generic value. Kegunaanya mirip dengan IF ELSE. Berikut untuk tipe Switch Yang ada :
	1. Single Switch Hanya menangani Tipe kondisi variabel
	2. Multiple Handler Switch Bias menangani beberapa tipe kondisi variabel dari setiap block kode
*/

package main

import "fmt"

func main() {
	// Single Switch
	hari := "senin"
	switch hari {
	case "senin":
		fmt.Println("Hari Ini Upacara")
	case "selasa":
		fmt.Println("Hari Ini Ngerjain Project 1")
	case "rabu":
		fmt.Println("Hari Ini Presentasi Project")
	case "kamis":
		fmt.Println("Hari Ini Libur")
	case "jumat":
		fmt.Println("Hari Ini Senam")
	default:
		fmt.Println("Hari Ini Libur!")
	}

	// Multiple Switch
	hari = "jumat"
	switch hari {
	case "jumat":
		fmt.Println("Hari Ini Senam!")
	default:
		fmt.Println("Ga Tau, Ga Mau Ngapain!")
	}
}
