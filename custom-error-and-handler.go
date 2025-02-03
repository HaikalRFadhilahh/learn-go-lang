/*
	CUSTOM ERROR meruapakan sebuah teknik membuat error secara manual dengan cara seperti ini user / developer memiliki kendali penuh akan implementasi error seperti apa. Tugas dari Developer yaitu membuat struct dan struct method untuk memenuhi syarat dari interface golang
*/

package main

import (
	"fmt"
)

/*
	Membuat Error Struct Not Found
*/

type ErrorNotFound struct {
	statusCode int
	message    string
}

func (e *ErrorNotFound) Error() string {
	return e.message
}

func (e *ErrorNotFound) StatusCode() int {
	return e.statusCode
}

/*
	Membuat Error Struct Internal Server Error
*/

type ErrorInternalServerError struct {
	statusCode int
	message    string
}

func (e *ErrorInternalServerError) Error() string {
	return e.message
}

func (e *ErrorInternalServerError) StatusCode() int {
	return e.statusCode
}

func main() {
	// Membuat Error
	var err error = &ErrorNotFound{
		message:    "Route Not Found",
		statusCode: 404,
	}

	// Membuat Case
	switch err.(type) {
	case *ErrorNotFound:
		fmt.Println("Error Karena Not Found (404)")
	case *ErrorInternalServerError:
		fmt.Println("Error Karena Internal Server Error (500)")
	default:
		fmt.Println("Error Not Identified!")
	}

}
