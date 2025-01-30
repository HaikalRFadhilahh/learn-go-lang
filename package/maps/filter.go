/*
	PACKAGE Go-Lang adalah merupakan sebuah Fungsi yang dapat di buat sendiri ataupun secara Public. Pada dasarnya PACKAGES adalah sebuah fungsi didalam File Golang yang dapat di gunakan oleh package lain maupun package utama ataupun main. Dalam Menggunakan sebuah Package kita dapat menbuat fungsi yang dapat di gunakan untuk tujuan spesifik dan memudahkan kita dalam Membuat Sebuah Aplikasi. Pastikan Saat kita membuat package Di Golang file packages di simpan pada folder terpisah agar memudahkan dalam manage
	! Perhatikan Struktur Folder Dalam Package Saat Membuat Sebuah Packages
	! Nama Package Pada Setiap File Bisa Menggunakan Nama Folder Saja
*/

package maps

/*
	Dalam File Ini Saya Akan Membuat Sebuah Function Package Untuk maps untuk memfilter data Yang Value Ingin Kita Ambil Berdasarkan Key Dari Slice Dan Akan Return Sebuah Slice Yang Akan Menampung Data Value Dari key Yang Diseleksi.
*/

/*
	Agar Package / Function yang di deklarasikan bisa digunakan lintas Package maka wajib menamainya dengan Huruf Kapital DiAwal
*/

func FilterMaps(data map[any]any, keyFilter []any) []any {
	temp := make([]any, 0)
	for _, v := range keyFilter {
		if data[v] != "" {
			temp = append(temp, data[v])
		}
	}

	return temp
}
