package internal

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func FILE() (fn string) {
	_, fn, _, _ = runtime.Caller(1)
	return filepath.Dir(fn)+string(os.PathSeparator)
}

func StringToInt(input string) int {
	i, _ :=strconv.Atoi(input)
	return i
}

func GetScannerForInput(inputFile string) *bufio.Scanner {
	input, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return  bufio.NewScanner(bufio.NewReader(input))
}