package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("vim-go")
}

func CreateZip() {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
}
