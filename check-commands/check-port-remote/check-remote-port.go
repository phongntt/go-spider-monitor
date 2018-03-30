package main

import (
  //"fmt"
  "github.com/reiver/go-telnet"
)

func SetTest() {
  conn, err := telnet.DialTo("192.168.72.123:8080")
  //conn.Write([]byte("hello world"))
  //conn.Write([]byte("\n"))
  if err != nil {
    println("ERROR")
    panic(err)
  }

  conn.Close()
  println("OK")
}

func main() {
  SetTest()
}