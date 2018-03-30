package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
  fileName := os.Args[1]
  strChekDuration := os.Args[2]

  checkDuration, err := time.ParseDuration(strChekDuration)
  if err != nil {
		fmt.Println("Argument 'duration' is not valid.")
		os.Exit(1)
  }

	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("Can't get file status.")
		os.Exit(1)
	}
	
	n := time.Now();

	fmt.Println("Now: ", n)
	fmt.Println("File time: ", info.ModTime())

	duration := n.Sub(info.ModTime())
	fmt.Println("Duration: ", duration)
	fmt.Println("Check Duration: ", checkDuration)

	//if n.After(info.ModTime().Add(checkDuration*time.Second)) {
	if n.After(info.ModTime().Add(checkDuration)) {
		fmt.Println("File have not been edited for too long.")
		os.Exit(1)
	}

	fmt.Println("File duration OK.")
	os.Exit(0)
}