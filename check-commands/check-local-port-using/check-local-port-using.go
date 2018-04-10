package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		//fmt.Fprintf(os.Stderr, "1- WARNING| Input port is missing.")
		fmt.Print("3- UNKNOWN| Input port is missing.")
		os.Exit(3)
	}

	port := args[0]
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		print(fmt.Sprintln(err))
		fmt.Print("3- UNKNOWN| Invalid port param.")
		os.Exit(3)
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		print(fmt.Sprintln(err))
		fmt.Print("0- OK| Port was using.")
		os.Exit(0)
	}

	ln.Close()
	fmt.Printf("1- WARNING| Port is available.")
	os.Exit(1)
}
