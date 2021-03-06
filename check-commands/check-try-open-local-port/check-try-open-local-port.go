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
		fmt.Print("1- WARNING| Input port is missing.")
		println()
		os.Exit(1)
	}

	port := args[0]
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "1- WARNING| Invalid port.")
		fmt.Print("1- WARNING| Invalid port.")
		println()
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		//fmt.Fprintf(os.Stderr, "1- WARNING| Can't listen on port.")
		fmt.Print("1- WARNING| Can't listen on port.")
		println()
		os.Exit(1)
	}

	err = ln.Close()
	if err != nil {
		//fmt.Fprintf(os.Stderr, "0- OK| Couldn't stop listening on port.")
		fmt.Print("0- OK| Can listening port but couldn't stop listening on it.")
		println()
		os.Exit(0)
	}

	fmt.Print("0- OK| TCP Port is available.")
	println()
	os.Exit(0)
}
