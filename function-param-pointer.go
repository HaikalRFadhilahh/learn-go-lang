/*
FUNCTION PARAMATER PASS POINTER meruapakan sebuah teknik mengirimkan sebuah alamat / pointer ke dalam paramater function, Yang sudah dijelaskan sebelumnya bahwa replikasi variabel / paramater function menggunakan Pass By Value sehingga dalam function ketika mengirimkan sebuah value variabel yang dirubah dalam function tidak akan mempengaruhi variabel sebelumnya. Ada sebuah kondisi atau kasus yang mengharuskan mengubah value variabel luar dengan function. Maka dari itu sebuah Parameter Pointer diperlukan dalam kasus ini.
*/

/*
	! Berhati-Hati Dalam Menggunakan Tehnik Paramter Pointer karena akan merubah value asli dari variabel diluar function, Gunakan dengan baik dan benar
*/

/*
	! IMPORTANT #1
	jika ingin sebuah function menerima alamat / pointer gunakan cara 'nama_variabel_pointer *tipe_data' dan untuk mengirimkannya gunakan '&nama_variabel_non_pointer' namun ketika variabel sudah pointer bisa langsung 'nama_variabel_pointer'

	! IMPORTANT #2
	Paramater menerima dalam bentuk alamat / pointer jika ingin mengambil value dari function bisa menggunakan * untuk pemgambilan value
*/

package main

import "fmt"

/*
	Function Pass By Value Yaitu Menghitung Total Angka Dari Slice
*/

func subNumberPBV(totalValue int, nomor ...int) {
	// Set Total Value To
	totalValue = 0

	// Sum Variadic Param Numbers
	for _, v := range nomor {
		totalValue += v
	}
}

/*
	Function Pass By Reference Yaitu Menghitung Total Angka Dari Slice
*/

func sumNumberPBR(totalValue *int, nomor ...int) {
	// Set Total Value To
	*totalValue = 0

	// Sum Variadic Param Numbers
	for _, v := range nomor {
		*totalValue += v
	}
}

func main() {
	// Declare Variabel
	totalValuePBV := 0
	totalValuePBR := 0

	// Print Default Variabel Value
	fmt.Println("Nilai Default Dari totalValue Pass By Value Adalah : ", totalValuePBV)
	fmt.Println("Nilai Default Dari totalValue Pass By Reference Adalah : ", totalValuePBR)

	// Break
	fmt.Println()

	// Execc Function
	subNumberPBV(totalValuePBV, 10, 20, 30, 40)
	sumNumberPBR(&totalValuePBR, 10, 20, 30, 40)

	// Print Value After Passing Variabel
	fmt.Println("Nilai Setelah Exec Function totalValue Pass By Value Adalah : ", totalValuePBV)
	fmt.Println("Nilai Setelah Exec Function totalValue Pass By Reference Adalah : ", totalValuePBR)
}
