package main

import (
	"AdventOfCode2021/internal"
	"fmt"
	"strconv"
)


func main() {
	scanner := internal.GetScannerForInput(internal.FILE()+"input")
	scanner.Scan()
	firstReading,_ := strconv.Atoi(scanner.Text())
	totalIncrease := 0
	for scanner.Scan() {
		nextReading,_ := strconv.Atoi(scanner.Text())
		fmt.Printf("%d > %d?\n",nextReading,firstReading)
		if nextReading > firstReading {
			totalIncrease += 1
		}
		firstReading = nextReading
	}
	fmt.Println(totalIncrease)
}
