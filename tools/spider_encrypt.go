package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/phongntt/crypto_helper"
)

func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	inputFilename := "input.txt"
	outputFilename := "output.txt"
	if len(os.Args) >= 3 {
		inputFilename = os.Args[1]
		outputFilename = os.Args[2]
	}

	infile, err := os.Open(inputFilename)
	errPanic(err)
	defer infile.Close()

	outfile, erro := os.Create(outputFilename)
	errPanic(erro)
	defer outfile.Close()

	scanner := bufio.NewScanner(infile)

	//1st line is the key
	if !scanner.Scan() {
		fmt.Println("Error when get the key.")
		errPanic(scanner.Err())
	}
	key := scanner.Text()
	if len([]byte(key)) != 32 {
		panic("Key must be 32bytes length!")
	}

	for scanner.Scan() {
		fmt.Println("Read --> ", scanner.Text())

		textLine := scanner.Text()
		encTextLine, erre := crypto_helper.Encrypt_Base64(textLine, key)
		errPanic(erre)

		nwrite, errw := outfile.WriteString(encTextLine + "\n")
		errPanic(errw)
		fmt.Println("Write --> ", nwrite, "bytes")
	}
	outfile.Sync()
}
