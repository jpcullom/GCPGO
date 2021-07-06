package main

import (
	"os"
)

//FileFill fills the contents of a file with monkey time
func FileFill(filename string) error {
	f, err := os.Create("monkey.txt")
	if err != nil {
		panic("File not opened")
	}
	// remember to close the file
	defer f.Close()

	for i := 0; i < 10000; i++ {
		f.WriteString("MONKEY TIME!!!\n")
	}
	return nil
}
