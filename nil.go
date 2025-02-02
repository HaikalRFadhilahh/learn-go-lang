/*
NIL merupakan sebuah value yang menandakan bahwa suatu variabel tidak meyimpan address / lokasi memori apapun. Pada dasarnya berbeda dengan tipe_data klasik lainya bahwa mereka akan melakukan asign default value dan membiatkan sebuah variabel mendapatkan alamat memori dari default value. Ketika kita mendeklarasikan sebuah pointer, interface, slice, map maka akan default valuenya bukanlah value kosong namun nil. Dan saat mengisi sebuah variabel dari nil harus dengan bantuan 'new' ataupun mengisinya dengan address dengan tanda & atau reference. Berikut Jenis-Jenis Tipe Data Lanjutan yang memiliki Default value nil :
1. Interface (Default value Nya adalah Nil [Namun Address Nya Sudah Ada])
2. Map (Default Value Nya adalah map[] namun itu sama dengan nil [namun addressnya juga sudah ada])
3. Slice (Default Valuenya adalah [] / nil. [Address Nya Juga Sudah Di Alokasikan])
4. Pointer
*/
package main

import "fmt"

/*
	Membuat Interface untuk Demo Nil Variabel
*/

type mahlukHidup interface {
	bernafas()
	jenisMahlukHidup() string
}

/*
	Membuat Struct dan Method Struct Untuk Demo Interface Implementation
*/

type manusia struct {
	nama  string
	jenis string
}

func (m *manusia) bernafas() {
	fmt.Printf("%v Sedang Bernafas...\n", m.nama)
}

func (m *manusia) jenisMahlukHidup() string {
	return m.jenis
}

func main() {
	// Demo Interface
	var haikal mahlukHidup

	fmt.Printf("Value Interface Default variabel haikal adalah : %v\n", haikal)
	fmt.Printf("Isi Reference Interface Default Value haikal adalah : %v\n", &haikal)

	// Mengisi Variabel Interface Dengan Impl Struct
	haikal = &manusia{
		nama:  "Haikal Raditya Fadhilah",
		jenis: "Manusia",
	}

	// Print Ulang Interface Value
	fmt.Printf("Value Interface Setelah Di Isi Impl Struct : %v \n", haikal)
	fmt.Printf("Isi Reference Interface Setelah Di Isi Value haikal adalah : %v\n", &haikal)

	// Break
	fmt.Println()

	// Demo Map
	var laptop map[string]string

	fmt.Printf("Isi Value Variabel Laptop Default Value Adalah :%v \n", laptop)
	fmt.Printf("Isi Reference Variabel Laptop Default Value Adalah :%v \n", &laptop)

	// Isi Map
	laptop = map[string]string{"Apple": "Macbook"}

	// Print Ulang
	fmt.Printf("Isi Variabel Laptop Setelah Di Isi Adalah : %v\n", laptop)
	fmt.Printf("Isi Reference Variabel Laptop Setelah Diisi Adalah :%v \n", &laptop)
	// Break
	fmt.Println()

	// Demo Slice
	var namaMahasiswa []string

	fmt.Printf("Isi Value Default Slice Kosong : %v\n", namaMahasiswa)
	fmt.Printf("Isi Reference Default Slice Kosong : %v\n", &namaMahasiswa)

	// Isi Slice
	namaMahasiswa = append(namaMahasiswa, "Haikal", "Vania", "Fajar")

	// Print Slice Ulang
	fmt.Printf("Isi Value Setelah Di Isi Slice  : %v\n", namaMahasiswa)
	fmt.Printf("Isi Reference Setelah Di Isi Slice : %v\n", &namaMahasiswa)

	fmt.Println("")

	// Break
	fmt.Println()

	// Demo Pointer
	var ptrAsalUniversitas *string

	// Print Default Value Pointer Asal Universitas
	// ! Error Karena Yang Nil itu Addressnya Bukan Value Nya. Dan ketika Reference Nil maka tidak bisa di dereference / extract valuenya.
	// fmt.Printf("Isi Value Variabel Pointer Kosongan : %v\n", *ptrAsalUniversitas) // Menggunakan * karena ptrAsalUniversitas Kita Menyimpan Memory nya
	// [Tidak Error]
	fmt.Printf("Isi Reference Variabel Pointer Kosongan : %v\n", ptrAsalUniversitas)

	// Isi Variabel Pointer [Asign Nil Reference Var Pointer Bisa dengan 2 Cara]
	/* MENGGUNAKAN VALUE  */
	ptrAsalUniversitas = new(string)
	*ptrAsalUniversitas = "Universitas Amikom Yogyakarta"

	/* MENGGUNAKAN REFERENCE (BUTUH VARIABEl LAIN) */
	var tmpAsalUniversitas string = "Universitas Bina Nusantara"
	ptrAsalUniversitas = &tmpAsalUniversitas

	// Print Ulang Value Pointer Setelah Asign
	fmt.Printf("Isi Reference Variabel Pointer Setelah Di Isi : %v\n", ptrAsalUniversitas)
	fmt.Printf("Isi Value Variabel Pointer Setelah Di Isi : %v\n", *ptrAsalUniversitas)
}
