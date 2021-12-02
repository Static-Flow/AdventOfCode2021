package main

import (
	"AdventOfCode2021/internal"
	"fmt"
	"strconv"
)

/*
https://adventofcode.com/2021/day/1
Solution for part 1
*/

func Version1() int {
	scanner := internal.GetScannerForInput(internal.FILE() + "input")
	//read the first line of input
	scanner.Scan()
	firstReading, _ := strconv.Atoi(scanner.Text())
	var nextReading int
	//tracks total increase changes
	totalIncrease := 0
	//for each line of input
	for scanner.Scan() {
		//convert it to an int
		nextReading, _ = strconv.Atoi(scanner.Text())
		//if it's greater than the previous input add to our sum
		if nextReading > firstReading {
			totalIncrease += 1
		}
		//set the previous tracked reading to our current reading
		firstReading = nextReading
	}
	return totalIncrease
}

func main() {
	//read in the input

	fmt.Println(Version1())
}
