/*
	Slice dalam Go meruapakan sebuah array yang fleksibel dapat menambah ukuran maupun mengecilkan ukuran. Dalam C++ Lebih dikenal dengan Vectors. Untuk membuat slice kita tidak perlu mendeklarasikan ukuran array. Kenapa Array? Karena Slice pada dasarnya adalah Array dengan Fitur Shink and Grow (Dapat Di Kecilkan dan di Besarkan) berbeda dengan array yang static ukurannya. Ada beberapa cara pembuatan slice sebagai berikut :
	1. var mySlice = []tipe_data{} / mySlice := []tipe_data{} = Dengan cara ini anda bisa membuat slice dari 0 dan dapat dimanipulasi dengan merubah ukuran atau menambah data
	2. var mySlice = myArr[index_awal:index_akhir] / mySlice := myArr[index_awal:index_akhir] = Dengan cara ini kita membuat Slice dengan cara memotong array yang sudah tersedia.
	3. var mySlice = make([]tipe_data, ukuran_slice_default, ukuran_maksimal_slice) / mySlice = make([]tipe_data, ukuran_slice_default, ukuran_maksimal_slice) = Dengan cara ini kita bisa menentukan batas maksimal slice dapat di Shrink dan Grow sehingga tidak Memenuhi Memori
*/

/*
	Extra Function Untuk Melihat Ukuran Default Dan Ukuran Maksinal Slice
	len(nama_variabel_slice / array) = Untuk Melihat Ukuran Slice / Array Default
	cap(nama_variabel_slice / array) = Untuk Melihat Ukuran Maksimal Dari Slice / Array ketika Di Konfigurasi
*/

package main

import "fmt"

func main() {
	// Membuat Slice Kosong
	sliceNamaMahasiswa := []string{}

	// Membuat Slice Dari Potongan Array
	arrNamaDaerah := [5]string{"Bantul", "Yogyakarta", "Sleman", "Gunung Kidul", "Kulon Progo"}
	sliceNamaDaerah := arrNamaDaerah[:len(arrNamaDaerah)-1]

	// Membuat Slice Dari Make
	sliceNamaUniversitas := make([]string, 1, 50)

	// Menampilkan Data Slice
	fmt.Println("Data Slice Nama Mahasiswa Dari Slice Kosong :", sliceNamaMahasiswa)
	fmt.Println("Ukuran Default Slice Kosong :", len(sliceNamaMahasiswa))
	fmt.Println("Ukuran Maksimal Slice Kosong :", cap(sliceNamaMahasiswa))

	fmt.Println("Data Slice Nama Daerah Dari Slice Potongan Array :", sliceNamaDaerah)
	fmt.Println("Ukuran Default Slice Potongan Array :", len(sliceNamaDaerah))
	fmt.Println("Ukuran Maksimal Slice Potongan Array :", cap(sliceNamaDaerah))

	fmt.Println("Data Slice Nama Daerah Dari Slice Fungsi Make :", sliceNamaUniversitas)
	fmt.Println("Ukuran Default Slice Fungsi Make :", len(sliceNamaUniversitas))
	fmt.Println("Ukuran Maksimal Slice Fungsi Make :", cap(sliceNamaUniversitas))

}
