package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./books.csv")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	readCSV(file)
}

func readCSV(file *os.File) {
	var r = csv.NewReader(file)
	for {
		var data, err = r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		for i := range data {
			fmt.Println(data[i])
		}
	}
}
