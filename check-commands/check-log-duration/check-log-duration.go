package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := os.Args[1]
	strChekDuration := os.Args[2] //ex: 1m, 1d, 2h30m2s

	checkDuration, err := time.ParseDuration(strChekDuration)
	if err != nil {
		fmt.Println("3- UNKNOWN| Argument 'duration' is not valid.")
		os.Exit(3)
	}

	info, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("3- UNKNOWN| Can't get log-file info.")
		os.Exit(3)
	}

	n := time.Now()

	println(fmt.Sprint("Now: ", n))
	println(fmt.Sprint("File time: ", info.ModTime()))

	duration := n.Sub(info.ModTime())
	println("Duration: ", duration)
	println("Check Duration: ", checkDuration)

	//if n.After(info.ModTime().Add(checkDuration*time.Second)) {
	if n.After(info.ModTime().Add(checkDuration)) {
		fmt.Println("1- WARNING| File have not been edited for too long.")
		os.Exit(1)
	}

	fmt.Println("0- OK| File duration OK.")
	os.Exit(0)
}
