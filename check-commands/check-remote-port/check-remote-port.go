package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		//fmt.Fprintf(os.Stderr, "3- UNKNOWN| Input arguments is missing.")
		fmt.Print("3- UNKNOWN| Input arguments is missing.")
		os.Exit(1)
	}

	port := args[1]
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "3- UNKNOWN| Invalid port %q: %s.", port, err)
		fmt.Print("3- UNKNOWN| Invalid 'port' param.")
		os.Exit(1)
	}

	remoteIp := args[0]

	// TimeOut in 10s
	conn, err := net.DialTimeout("tcp", remoteIp+":"+string(port), 10*time.Second)
	if err != nil {
		//fmt.Fprintf(os.Stderr, "1- WARNING| Cannot connect to remote port.")
		fmt.Print("1- WARNING| Cannot connect to remote port.")
		os.Exit(1)
	}
	defer conn.Close()

	//fmt.Fprintf(os.Stderr, "0- OK| Success telneting to remote port.")
	fmt.Print("0- OK| Success telneting to remote port.")
	os.Exit(0)
}
