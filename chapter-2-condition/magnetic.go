package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filePath, err := filepath.Abs("my.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err2 := os.Stat(filePath)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(fileInfo.Size())
}
