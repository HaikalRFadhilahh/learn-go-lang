/*
	TYPE-INFERENCE dalam Go-Lang Merupakan Salah satu cara deklarasi Variabel Yang memungkinkan go-lang mendeklarasikan sebuah tipe data tanpa mendeklarasikan secara eksplisit. Namun ada persyaratan yang harus di penuhi seperti harus ada data yang dimasukkan ke variabel, Bisa dari data langsung atau return dari hasil function. Atau untuk menyimpan sebuah Anonymous Function yang nanti ketika di gunakan adalah berbentuk function. Kita Bisa Mendekalarasikan Variabel Type Inference dengan cara 'nama_variabel := value / anonymous function'.
*/

package main

import "fmt"

/*
Membuat Function Untuk Simulasi Type Inference Dari Return Function
*/
func perkalian(a int, b int) int {
	return a * b
}

func main() {
	// Type Inferences Dengan Value Langsung (string)
	namaMahasiswa := "Haikal Raditya Fadhilah"

	// Type Inference Dengan Return Function (Function Perkalian)
	nilai1 := 10
	nilai2 := 20
	hasil := perkalian(nilai1, nilai2)

	// Type Inference Untuk Menyimpan Anonymous Function
	sayHello := func() {
		fmt.Println("Hello World!")
	}

	// Print Out
	fmt.Printf("Nama Mahasiswa Adalah : %v\n", namaMahasiswa)
	fmt.Printf("Hasil Perkalian '%v' x '%v' Adalah %v \n", nilai1, nilai2, hasil)
	sayHello()

}
