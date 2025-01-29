/*
	MAP / MAPS pada golang merupakan sebuah Tipe Data untuk membuat struktur data Key:Value mirip seperti Json ketika di Javascript. Dengan MAP kita dapat meyimpan data dengan key yang kita inginkan kita juga dapat mendeklarasikan tipe data kunci dan tipe data value pada MAP, Cara pembuatan MAP ada 2 cara yaitu :
	1. var nama_variabel map[tipe_data_key]tipe_data_value (Kurang Recomended) Akan Error Dan harus di make dulu
	2. nama_variabel := make(map[tipe_data_key]tipe_data_value) (Recomended ) Bisa Langsung Digunakan
	3. nama_variabel := map[tipe_data_key]tipe_data_value{} / var nama_variabel = map[tipe_data_key]tipe_data_value{} (Standar) Bisa Langsung Digunakan
*/

package main

import "fmt"

/*
DEKLARASI Alias Untuk MAP Go-Lang
*/
type mapKegiatanTahunan map[int]string
type mapMahasiswaDosenPembimbing map[string]string

func main() {
	// Membuat MAP dengan asign langsung data ke variabel
	mapKegiatan := mapKegiatanTahunan{2021: "Mahasiswa Baru", 2024: "Internship MSIB LLDIKTI 5 Yogyakarta", 2020: "Karantina Covid-19", 2022: "Asisten Praktikum Baru", 2023: "IT Support FA"}

	// Membuat MAP Kosong Dan Insert Data MAP Secara Manual
	var mapMhsDosbing = make(mapMahasiswaDosenPembimbing)
	mapMhsDosbing["Haikal"] = "Pak Ika"
	mapMhsDosbing["Ariza"] = "Pak Nuri"
	mapMhsDosbing["Zidan"] = "Bu Rumini"

	// Menampilkan Map Dengan For (Seharusnya Data Yang Di Tampilkan Tidak Sesuai Ukuran Deklarasi, Kenapa?)
	for k, v := range mapMhsDosbing {
		fmt.Printf("Mahasiswa Atas Nama '%v' Di Bimbing Oleh '%v'\n", k, v)
	}

	// Space
	fmt.Println()

	// Menampilkan Data Secara Berurutan
	orderMapKey := [...]int{2021, 2024, 2020, 2023}
	for _, v := range orderMapKey {
		fmt.Printf("Tahun '%v' Kegiatan Saya Adalah '%v'\n", v, mapKegiatan[v])
	}
}
