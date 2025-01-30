/*
PACKAGE - IMPORTS - EXPORTS merupakan suatu teknik bagaimana kita membuat fungsi yang dapat Reusable (Dapat Digunakan Ulang Oleh Package Lain) Pada File ini kita akan melakukan Import Dari File `package/maps/filter.go` dengan Nama Fungsi FilterMaps
*/
package main

/*
	Sebelum Menggunakan Library / Package Local Pastikan Untuk Melakukan Import Nama Package Dari Folder Package sehingga Fungsi Yang Bersifat Global Dapat DI Gunakan DI Package Lain.
*/

import (
	"github.com/HaikalRFadhilahh/learn-go-lang/package/maps"
	"github.com/HaikalRFadhilahh/learn-go-lang/package/slice"
)

/*
	Fungsi Utama, Main Function
*/

func main() {
	// Membuat Maps
	mapsIdentitasMahasiswa := map[any]any{"21.11.3910": "Haikal R Fadhilah", "21.11.3911": "Habib", "21.11.3909": "Bari", "21.11.3912": "Febrian", "21.11.3913": "Zidan", "21.11.3902": "Asyraf"}

	// Membuat Slice Untuk Filter Key (Data Nama Yang ingin Di Ambil NIM 21.11.3910 dan 21.11.3902)
	mapsFilterKey := []any{"21.11.3910", "21.11.3902"}

	// Memfilter Data
	dataMahasiswaFiltered := maps.FilterMaps(mapsIdentitasMahasiswa, mapsFilterKey)

	// Menampilkan Data Setelah Di Filter
	slice.Show(dataMahasiswaFiltered)
}
