/*
MANIPUTLASI SLICES GO-LANG
Sudah dibahas sebelumnya bahwa slice dapat di manipulasi dengan sangat fleksibel mulai dari ukuran slice, copy isi slice bahkan membuat slice dari array yang sudah tersedia. Berikut Contoh manipulasi slice.
1. Memperbesar Ukuran Slice dengan perintah append dengan cara menambah data kedalam slice.
exp: append(nama_slice,data1,data2,data_selanjutnya)
2. Menggabungkan Slice dengan Slice yang lain dengan append (append bisa untuk push value dan menggabungkan slice)
exp: append(nama_slice_1,nama_slice_2)
3. Melakukan Copy slice ke slice yang lain
*/
package main

import "fmt"

func main() {
	// 1. Memperbesar Ukuran Slice Dengan Append
	sliceNamaMahasiswa := []string{"Haikal", "Ariza", "Shera"}
	fmt.Println("Slice Nama Mahasiswa Sebelum Di Append :", sliceNamaMahasiswa)
	sliceNamaMahasiswa = append(sliceNamaMahasiswa, "Yoko")
	fmt.Println("Slice Nama Mahasiswa Setelah Di Append :", sliceNamaMahasiswa)

	// 2. Menggabungkan Slice
	sliceNamaDosen1 := []string{"Pak Ika"}
	sliceNamaDosen2 := []string{"Pak Ferian"}
	sliceNamaDosenGabungan := append(sliceNamaDosen1, sliceNamaDosen2...)
	fmt.Println("Slice Data Gabungan Dosen :", sliceNamaDosenGabungan)

	// 3. Melakukan Copy Slice
	sliceNamaDaerah := []string{"Bantul", "Yogyakarta", "Sleman"}
	sliceNamaDaerahCopy := make([]string, 2)
	copy(sliceNamaDaerahCopy, sliceNamaDaerah)
	fmt.Println("Data Slice Nama Derah Copy :", sliceNamaDaerahCopy)
}
