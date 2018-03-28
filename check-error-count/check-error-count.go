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
	/*
		if err != nil {
			fmt.Println("Can't convert arguments to int.")
			panic(err)
		}
	*/
	errPanic(err)

	fmt.Println("----- Process log file --")
	smallFile := createSmallFile(fileToRead, byteNum)

	// log like this
	// 06/03/2017 14:57:28,765 DEBUG: com.dtsc.helijobs.logger.debug - ERROR ON: AuthenticationCommand
	// --> reg = "^[0-9/]{10} [0-9/:,]{8,} (DEBUG|ERROR)?: .* ERROR ON:"
	errCount := readFileAndCountError(smallFile, "^[0-9/]{10} [0-9/:,]{8,} (DEBUG|ERROR)?: .* ERROR ON:")
	fmt.Println("----- End: Process log file --")
	fmt.Println("ErrCount: ", errCount)

	os.Remove(smallFile)

	if errCount > errLimit {
		fmt.Println("ErrCount is over limit!")
		os.Exit(1) // So loi vuot qua muc cho phep
	}

	fmt.Println("ErrCount is under limit!")
	os.Exit(0) // So loi nam trong gioi han cho phep
}

/*
Từ file lớn, tạo ra file nhỏ hơn với dung lượng bằng dung lượng được gửi từ tham số
*/
func errPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func createSmallFile(fname string, byteNum int) string {
	file, err := os.Open(fname)
	errPanic(err)
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
	errPanic(err)

	// Timestamp part in small_<Timestamp>.log is Unix Millisecond
	newfname := "small_" + strconv.FormatInt((time.Now().UnixNano()/1000000), 10) + ".log"
	err1 := ioutil.WriteFile(newfname, buf, 0644)
	errPanic(err1)

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
		fmt.Println("Can't convert string to int.")
		return "", -1, -1, err
	}

	errLimit, err := strconv.Atoi(errLimitStr)
	if err != nil {
		fmt.Println("Can't convert string to int.")
		return "", -1, -1, err
	}

	return fileName, byteNum, errLimit, nil
}

func readFileAndCountError(filename string, errRegex string) int {
	file, err := os.Open(filename)
	errPanic(err)
	defer file.Close()

	errCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		////fmt.Println(scanner.Text())
		if isErrLine(scanner.Text(), errRegex) {
			errCount++
		}
	}

	if err := scanner.Err(); err != nil {
		errPanic(err)
	}

	return errCount
}

func isErrLine(textLine, errRegex string) bool {
	regErrFilter := regexp.MustCompile(errRegex)
	isThisErrorLine := regErrFilter.MatchString(textLine)

	if isThisErrorLine {
		fmt.Println("---", textLine)
		fmt.Println("-----> THIS IS ERROR LINE <-----")
	} /*else {
		fmt.Println("-----> ------- <-----")
	}*/

	return isThisErrorLine
}
