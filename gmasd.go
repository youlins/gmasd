package main

import (
	"fmt"
	"./access_point"
)

const (
	VERSION = "0.0.1"
)

func main() {

	fmt.Printf("Hello gmasd %s\n", VERSION)

	ap := access_point.New("coderA")
	ap.Listen(8000)
	ap.Start()
}