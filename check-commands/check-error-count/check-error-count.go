package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	//fileToRead, byteNum, errLimit, err := readFromArgs()
	fileToRead, byteNum, errLimit, err := readFromArgs()
	if err != nil {
		fmt.Print("3- UNKNOWN| Arguments error")
		os.Exit(3)
	}

	println("----- Process log file --")
	smallFile := createSmallFile(fileToRead, byteNum)

	// log like this
	// 06/03/2017 14:57:28,765 DEBUG: com.dtsc.helijobs.logger.debug - ERROR ON: AuthenticationCommand
	// --> reg = "^[0-9/]{10} [0-9/:,]{8,} (DEBUG|ERROR)?: .* ERROR ON:"
	errCount := readFileAndCountError(smallFile, "^[0-9/]{10} [0-9/:,]{8,} (DEBUG|ERROR)?: .* ERROR ON:")
	println("----- End: Process log file --")
	println("ErrCount: ", errCount)

	os.Remove(smallFile)

	if errCount > errLimit {
		fmt.Print("1- WARNING| ErrCount is over limit.")
		os.Exit(1) // So loi vuot qua muc cho phep
	}

	fmt.Print("0- OK| ErrCount is under limit")
	os.Exit(0) // So loi nam trong gioi han cho phep
}

func createSmallFile(fname string, byteNum int) string {
	file, err := os.Open(fname)
	if err != nil {
		print(fmt.Sprintln(err))
		fmt.Print("3- UNKNOWN| Cannot open log file")
		os.Exit(3)
	}
	defer file.Close()

	stat, err := os.Stat(fname)
	readByteNum := int64(byteNum)
	start := stat.Size() - readByteNum
	if start < 0 {
		start = 0
		readByteNum = stat.Size()
	}

	buf := make([]byte, readByteNum)
	_, err = file.ReadAt(buf, start)
	if err != nil {
		print(fmt.Sprintln(err))
		fmt.Print("3- UNKNOWN| Cannot read log file")
		os.Exit(3)
	}

	// Timestamp part in small_<Timestamp>.log is Unix Millisecond
	newfname := "small_" + strconv.FormatInt((time.Now().UnixNano()/1000000), 10) + ".log"
	err1 := ioutil.WriteFile(newfname, buf, 0644)
	if err1 != nil {
		print(fmt.Sprintln(err1))
		fmt.Print("3- UNKNOWN| Cannot write a small log file")
		os.Exit(3)
	}

	return newfname
}

func readFromArgs() (string, int, int, error) {
	fileName := os.Args[1]
	byteNumStr := os.Args[2]
	errLimitStr := os.Args[3]
	println("[Args] File: ", fileName)
	println("[Args] Byte num: ", byteNumStr)
	println("[Args] Err limit: ", errLimitStr)

	byteNum, err := strconv.Atoi(byteNumStr)
	if err != nil {
		println("Can't convert string to int.")
		return "", -1, -1, err
	}

	errLimit, err := strconv.Atoi(errLimitStr)
	if err != nil {
		println("Can't convert string to int.")
		return "", -1, -1, err
	}

	return fileName, byteNum, errLimit, nil
}

func readFileAndCountError(filename string, errRegex string) int {
	file, err := os.Open(filename)
	if err != nil {
		print(fmt.Sprintln(err))
		fmt.Print("3- UNKNOWN| Cannot read small log file")
		os.Exit(3)
	}
	defer file.Close()

	/*
		errCount := 0
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			////fmt.Println(scanner.Text())
			if isErrLine(scanner.Text(), errRegex) {
				errCount++
			}
		}

		if err := scanner.Err(); err != nil {
			print(fmt.Sprintln(err))
			fmt.Print("3- UNKNOWN| Cannot read small log file")
			os.Exit(3)
		}
	*/

	println("Begin count err-line num")
	errCount := 0
	reader := bufio.NewReaderSize(file, 500000)
	for byteLine, _, err := reader.ReadLine(); err == nil; byteLine, _, err = reader.ReadLine() {
		////println(string(byteLine))
		if isErrLine(string(byteLine), errRegex) {
			errCount++
		}
	}

	if err != nil {
		print(fmt.Sprintln(err))
		fmt.Print("3- UNKNOWN| Cannot read small log file")
		os.Exit(3)
	}

	return errCount
}

func isErrLine(textLine, errRegex string) bool {
	regErrFilter := regexp.MustCompile(errRegex)
	isThisErrorLine := regErrFilter.MatchString(textLine)

	if isThisErrorLine {
		println("---", textLine)
		println("-----> THIS IS ERROR LINE <-----")
	} /*else {
		fmt.Println("-----> ------- <-----")
	}*/

	return isThisErrorLine
}
