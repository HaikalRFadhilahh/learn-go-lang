/*
	MODULES / Package Luar / Package Eksternal merupakan sebuah teknik dalam Envronment Golang untuk melakukan instalasi package dalam golang. Ada Beberapa perintah yang bisa digunakan untuk mendownload library / package, berbeda dengan bahasa pemrograman yang lain bahwa mereka membutuhkan library lain seperti Composer untuk PHP, Pip untuk Pyhton, NPM untuk Node js berbeda dengan golang yang memiliki perintah untuk mendownload library. Library Dalam golang hanya menggunakan sebuah Url Redirect yang menuju ke sebuah repository Yang dapat digunakanpada golang. berikut perintah pendukung Deployment Yang Dapat Digunakan :
	1. go mod (init, tidy, download) Init untuk membuat nama project golang, tidy untuk membersihkan cache dan download dependencies library, download untuk mengunduh semua library yang ada pada file go.mod
	2. go get (-u) url_library untuk mendownload Libray ketika menggunakan -u untuk melakukan update semua versi library / jika tanpa param -u hanya install tanpa update library versi lainnya
	3. go build -o lokasi_binary file.go Untuk melakukan build menjadi binary
	4. go test untuk melakukan Test

*/

/*
	MODULES CASES
	! Dalam File Ini Kita Akan Menggunakan Package mux dari url github.com/gorilla/mux@latest untuk membuat http server
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	Membuat Struct Untuk Response Json
*/

type SuccessResponse struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

/*
	FUNCTION Handler
*/

func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SuccessResponse{
		StatusCode: 200,
		Message:    "Go-Lang Application Programming Interface",
	})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&SuccessResponse{
		StatusCode: 404,
		Message:    "Not Found",
	})
}

/*
MAIN FUNCTION
*/
func main() {
	// Declare Router
	r := mux.NewRouter()

	// Route
	r.HandleFunc("/", Index)

	// Handler Not Found
	r.NotFoundHandler = http.HandlerFunc(NotFound)

	// Running
	fmt.Println("Server Starting...")
	if err := http.ListenAndServe("127.0.0.1:8000", r); err != nil {
		fmt.Println("Server Error :", err.Error())
	}

}
