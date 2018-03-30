package main

import (
  "net"
  "fmt"
  "flag"
  "os"
  "strconv"
)

func usage() {
    fmt.Fprintf(os.Stderr, "usage: test-check-local-port [port number]\n")
    flag.PrintDefaults()
    os.Exit(2)
}

func main() {
  flag.Usage = usage
  flag.Parse()

  args := flag.Args()
  if len(args) < 1 {
      fmt.Fprintf(os.Stderr, "Input port is missing.")
      os.Exit(1)
  }

  port := args[0]
  _, err := strconv.ParseUint(port, 10, 16)
  if err != nil {
    fmt.Fprintf(os.Stderr, "Invalid port %q: %s\n", port, err)
    os.Exit(1)
  }

  ln, err := net.Listen("tcp", ":" + port)

  if err != nil {
    fmt.Fprintf(os.Stderr, "Can't listen on port %q: %s\n", port, err)
    os.Exit(1)
  }

  err = ln.Close()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Couldn't stop listening on port %q: %s\n", port, err)
    os.Exit(0)
  }

  fmt.Printf("TCP Port %q is available\n", port)
  os.Exit(0)
}