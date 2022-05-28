package main

import (
	"bufio"
	"fmt"
	"image/png"
	"os"
	"strconv"

	"golang.org/x/image/webp"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	//take files from the command line
	inputFileList := os.Args[1:]
	totalFiles := len(inputFileList)
	currentPosition := 0

	//loop through all files
	for _, path := range inputFileList {
		currentPosition++
		progress := strconv.Itoa(currentPosition) + "/" + strconv.Itoa(totalFiles) + " "
		if path[len(path)-5:] == ".webp" {
			fmt.Println(progress + "Processing file: " + path)

			//convert webp to png
			source, err := os.Open(path)
			checkErr(err)

			defer source.Close()
			content, err := webp.Decode(source)
			checkErr(err)

			outputFile, err := os.Create(path[:len(path)-5] + ".png")
			checkErr(err)

			err = png.Encode(outputFile, content)
			checkErr(err)
		} else {
			fmt.Println(progress + "Skipping: " + path)
		}
	}

	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
