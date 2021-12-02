package main

import (
	"AdventOfCode2021/internal"
	"fmt"
)

func sumSweeps(sweeps []int) int {
	return sweeps[0] + sweeps[1] + sweeps[2]
}

/*
https://adventofcode.com/2021/day/1
Solution for part 2
*/

/*
This version scans in all values into memory then walks the values to compute the triplets.
It is 2.7x slower than Version1() but uses 2.5x less space

cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkVersion2-8   	   10000	    117824 ns/op
*/
func Version2() int {
	var sweepValues []int
	sweepIncreases := 0
	scanner := internal.GetScannerForInput(internal.FILE() + "input")
	for scanner.Scan() {
		sweepValues = append(sweepValues, internal.StringToInt(scanner.Text()))
	}
	sweepIndex := 0
	for i := 1; i < len(sweepValues)-2; i++ {
		previousSweepSum := sweepValues[sweepIndex+0] + sweepValues[sweepIndex+1] + sweepValues[sweepIndex+2]
		currentSweepSum := sweepValues[i+0] + sweepValues[i+1] + sweepValues[i+2]
		if previousSweepSum < currentSweepSum {
			sweepIncreases++
		}
		sweepIndex++
	}
	return sweepIncreases
}

/*
This version scans values one at a time and creates a length 3 array to hold it followed by back filling the value into the previous two triplet arrays.
After the back fill it checks if the previous two
It is 2.7x faster than Version2() but uses 2.5x more space

cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz
BenchmarkVersion1-8   	    3685	    303066 ns/op
*/
func Version1() int {
	//track our sweep pair inputs as an array of arrays containing the sweep triples
	var sweeps [][]int
	//track number of increasing sweep pairs
	sweepIncreases := 0
	//read the input
	scanner := internal.GetScannerForInput(internal.FILE() + "input")
	//The sweep index
	currentSweepIndex := 0
	//for each line of input
	for scanner.Scan() {
		//convert it to an int
		nextSweepValue := internal.StringToInt(scanner.Text())
		//make a new sweep triplet for the value and add it to its triplet along with back filling it to the previous 2
		sweeps = append(sweeps, []int{})
		for i := currentSweepIndex; i >= currentSweepIndex-2; i-- {
			sweeps[i] = append(sweeps[i], nextSweepValue)
			//if we get to the first triplet when back filling we stop
			if i == 0 {
				break
			}
		}
		//Once we have 4 sweep triplets we must have 2 that are full so we compare them
		if len(sweeps) > 3 {
			// if there is an increase between the triplets
			if sumSweeps(sweeps[currentSweepIndex-3]) < sumSweeps(sweeps[currentSweepIndex-2]) {
				//add one to the sweepIncrease
				sweepIncreases++
			}
		}
		//advance the sweep index
		currentSweepIndex += 1
	}
	return sweepIncreases
}

func main() {
	fmt.Println(Version1())
}
