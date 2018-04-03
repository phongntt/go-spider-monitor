package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: check-local-port [port number]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "1- FAIL| Input port is missing.\n")
		os.Exit(1)
	}

	port := args[0]
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		fmt.Fprintf(os.Stderr, "1- FAIL| Invalid port %q: %s\n", port, err)
		os.Exit(1)
	}

	ln, err := net.Listen("tcp", ":"+port)

	if err != nil {
		fmt.Fprintf(os.Stderr, "1- FAIL| Can't listen on port %q: %s\n", port, err)
		os.Exit(1)
	}

	err = ln.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "0- SUCCESS| Couldn't stop listening on port %q: %s\n", port, err)
		os.Exit(0)
	}

	fmt.Printf("0- SUCCESS| TCP Port %q is available\n", port)
	os.Exit(0)
}
