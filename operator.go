/*
OPERATOR meruapakan teknik untuk melakukan manipulasi nilai dan komprasi dari 2 nilai atau lebih mulai dari Nilai angka atau String dengan Operasi Perbandingan.
*/
package main

import "fmt"

/*
	Operasi Aritmatika Merupakan teknik untuk melakukan perhitungan atau operasi yang berhubungan nilainya angka. Tidak bisa di lakukan pada Tipe Data Lain Selain Angka. Berikut Operasi dan fungsinya yang dapat di lakukan :
	1. '+' Operasi ini untuk menambahkan 2 atau lebih nilai angka
	2. '-' Operasi untuk mengurangi nilai dari nilai lainnya
	3. '*' Operasi untuk melakukan perkalian nilai angka
	4. '/' Operasi Untuk Melakukan pembagian
	5. '%' Operasi Untuk Melakukan Modulus (Sisa Bagi)
	6. '++' Operasi Untuk melakukan Increment atau Penambahan Nilai + 1
	7. '--' Operasi Untuk melakukan Decrement atau Pengurangan Nilai - 1
*/

func aritmatika() {
	// Penambahan
	fmt.Println("10 + 5 Adalah :", 10+5)
	// Pengurangan
	fmt.Println("10 - 5 Adalah :", 10-5)
	// Perkalian
	fmt.Println("10 * 5 Adalah :", 10*5)
	// Pembagian
	fmt.Println("10 / 2 Adalah :", 10/5)
	// Modulus
	fmt.Println("10 % 3 Adalah :", 10%3)
	// Increment
	i := 10
	i++
	fmt.Println("Increment dari 10 :", i)
	// Decrement
	d := 10
	d--
	fmt.Println("Decrement Dari 10 :", d)
}

/*
	OPERATOR BITWISE Merupakan Operator Untuk Binary 0 dan 1. Yang nanti akan di convert ke angka kembali. Jadi ketika melakukan komparasi dari angka, maka di convert ke biner lalu di lakukan operasi. Setelah success maka di konvert kembali ke angka / tipe data sebelumnya. Berikut Daftar Operasi BITWISE :
	1. '&' / AND
	2. '|' / OR
	3. '^' / XOR
	4. '<<' / Shift Biner Left
	5. '>>' / Shift Biner Right
*/

func bitwise() {
	angka1 := 10 // 1010 (biner)
	angka2 := 4  // 0100 (biner)

	fmt.Printf("Biner dari %d: %b\n", angka1, angka1)
	fmt.Printf("Biner dari %d: %b\n", angka2, angka2)

	// 1. '&' (AND)
	fmt.Printf("\nHasil AND (%d & %d): %d, Biner: %b\n", angka1, angka2, angka1&angka2, angka1&angka2)

	// 2. '|' (OR)
	fmt.Printf("Hasil OR (%d | %d): %d, Biner: %b\n", angka1, angka2, angka1|angka2, angka1|angka2)

	// 3. '^' (XOR)
	fmt.Printf("Hasil XOR (%d ^ %d): %d, Biner: %b\n", angka1, angka2, angka1^angka2, angka1^angka2)

	// 4. '<<' (Shift Biner Left)
	fmt.Printf("\nHasil Shift Left (%d << 1): %d, Biner: %b\n", angka1, angka1<<1, angka1<<1)
	fmt.Printf("Hasil Shift Left (%d << 2): %d, Biner: %b\n", angka1, angka1<<2, angka1<<2)

	// 5. '>>' (Shift Biner Right)
	fmt.Printf("\nHasil Shift Right (%d >> 1): %d, Biner: %b\n", angka1, angka1>>1, angka1>>1)
	fmt.Printf("Hasil Shift Right (%d >> 2): %d, Biner: %b\n", angka1, angka1>>2, angka1>>2)
}

/*
	OPERATOR KOMPARASI adalah operator untuk melakukan perbandingan dari 2 value atau lebih. Bisa di gunakan pada value angka maupun string tapi tetap tergantung dari kondisinya. Berikut jenis-jenis operator komparasi :
	1. '==' Sama Dengan Akan return True jika nilainya sama
	2. '!=' Tidak Sama Dengan akan return True jika Nilainya Berbeda
	3. '>' Lebih Besar Akan Return True Jika Nilai Pertama Lebih Besar Daripada Nilai Kedua
	4. '<' Lebih Kecil Akan Return True Jika Nilai Pertama Lebih Kecil Daripada Nilai Kedua
	5. '>=' Lebih Besar Sama Dengan Akan Return True Jika Nilai Lebih besar atau Sama Dengan Nilai Kedua
	6. '<=' Lebih Kecil Sama Dengan Akan Return True Jika Nilai Pertama Lebih Kecil Atau Sama Dengan Nilai kedua
*/

func komprasi() {
	a := 530
	b := 400

	// Sama Dengan
	fmt.Println("Apakah 503 Sama Dengan 400? :", a == b)
	// Tidak Sama Dengan
	fmt.Println("Apakah 503 Tidak Sama Dengan 400? :", a != b)
	// Lebih Besar
	fmt.Println("Apakah 503 Lebih Besar Daripada 400? :", a > b)
	// Lebih Kecil
	fmt.Println("Apakah 503 Lebih Kecil Daripada 400? :", a < b)
	// Lebih Besar Sama Dengan
	fmt.Println("Apakah 503 Lebih Besar Sama Dengan 400? :", a >= b)
	// Lebih Kecil Sama Dengan
	fmt.Println("Apakah 503 Lebih Kecil Sama Dengan 400? :", a <= b)
}

/*
	OPERATOR LOGICAL meruapakn Operator untuk operasi perbandingan antara 2 atau lebih nilai Boolean entah dari boolean atau dari operasi KOmparasi dari Nilai Lain Non Booleaan. Berikut Daftar Operator LOGICAL :
	1. '&&' AND = Akan return nilai True jika kedua kondisi Nilainya True
	2. '||' OR = Akan Return Nilai True Jika Salah Satu Komparasi Nilainya True
	3. '!' NOT = Akan Membalikkan Nilai Boolean Sebelumnya. Contoh : Akan Nilai True jika !false atau !nama_variabel_nilai_false
*/

func logical() {
	bool1, bool2 := true, false

	// AND
	fmt.Println("Logical AND Nilai True Dan False Adalah :", bool1 && bool2)
	// OR
	fmt.Println("Logical OR Nilai True Dan False Adalah :", bool1 || bool2)
	// NOT
	fmt.Println("Logical NOT Nilai True Adalah :", !bool1)
}

/*
	OPERATOR ASSIGNMENT Adalah Operator yang mirip seperti aritmatika namun hasil dari OPERASI dapat langsung di asign atau di masukkan ke variabel. Berikut Daftar Opertor ASSIGNMENT :
	1. '=' Opeator Ini akan melakukan deklarasi dan masukkan nilai ke variabel
	2. '+=' Operator ini akan menambahkan Nilai Sebelumnya nilai baru lalu di masukkan ke variabel
	3. '-=' Operator Ini Akan Mengurangi Nilai Variabel dengan nilai tertantu dan perbaharui ke variabel
	4. '*=' Operator ini akan melakukan perkalian nilai Variabel dengan nilai baru dan asign ke variabel
	5. '/=' Operator ini akan melakukan pembagian nilai Variabel dengan nilai baru dan asign ke variabel
	6. '%=' Operator ini akan melakukan modulus nilai Variabel dengan nilai baru dan asign ke variabel
	7. '&=' Operator Akan Melakukan Bitwise AND dan asign Ke Variabel
	8. '|=' Operator Akan Melakukan Bitwise OR dan Asign ke Variabel
	9. '^=' Operator Akan Melakukan Bitwise XOR dan Asign ke Variabel
	10. '>>=' Operator Akan Melakukan Bitwise Shift Biner Ke Kanan dan Asign ke Variabel
	11. '<<=' Operator Akan Melakukan Bitwise Shift Biner Ke Kiri dan Asign ke Variabel
*/

func assignment() {
	n := 5

	// Operator =
	variabel := 100
	fmt.Println("Nilai Varibel di asign nilai 100, Berikut Outpunya :", variabel)

	// Operator +=
	variabel += n
	fmt.Println("Nilai Variabel Setelah Menggunakan Operator += Adalah :", variabel)
	variabel = 100

	// Operator -=
	variabel -= n
	fmt.Println("Nilai Variabel Setelah Menggunakan Operator -= Adalah :", variabel)
	variabel = 100

	// Operator *=
	variabel *= n
	fmt.Println("Nilai Variabel Setelah Menggunakan Operator *= Adalah :", variabel)
	variabel = 100

	// Operator /=
	variabel /= n
	fmt.Println("Nilai Variabel Setelah Menggunakan Operator /= Adalah :", variabel)
	variabel = 100

	// Operator %=
	variabel %= n
	fmt.Println("Nilai Variabel Setelah Menggunakan Operator %= Adalah :", variabel)
	variabel = 100
}

func main() {
	assignment()
}
