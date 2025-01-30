/*
	PACKAGE SLICE Go-Lang Merupakan Sebuah Simulasi untuk mendemonstrasikan Package Dalam Go-Lang Dan Reusable Function Yang Dalam Menampilkan Seluruh Data Slice
*/

package slice

import "fmt"

func Show(data []any) {
	fmt.Println("Data Dari Slice : ")
	for i, v := range data {
		fmt.Printf("%v. %v\n", i+1, v)
	}
}
