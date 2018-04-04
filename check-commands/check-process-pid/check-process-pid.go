package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func findProcess(pid int) (int, string) {
	println("----------- Begin findProcess ----------------")
	app := "/bin/ps"
	arg0 := "--pid=" + strconv.Itoa(pid)

	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()

	if err != nil {
		println("----------- End findProcess ----------------")
		return 1, "Can't run PS command."
	}

	// XU LY STD OUTPUT DE LAY THONG TIN
	//----------------------------------
	strOutput := string(stdout)
	println(strOutput)

	println("----------- End findProcess ----------------")
	return 0, "Success"
}

func main() {
	flag.Parse()
	s := flag.Arg(0)

	//processId := 3187
	processId, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("1- WARNING| Invalid PID arg.")
		os.Exit(1)
	}

	println("Find PID: ", processId)

	isSucc, _ := findProcess(processId)
	//fmt.Println(isSucc)
	//fmt.Println(errMsg)

	if isSucc != 0 {
		fmt.Println("1- WARNING| Can't find process.")
		os.Exit(1)
	}

	fmt.Println("0- SUCCESS| Process is exists.")
	os.Exit(0)
}
