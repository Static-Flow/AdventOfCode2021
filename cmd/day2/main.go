package main

import (
	"AdventOfCode2021/internal"
	"fmt"
)

func zeroOrIMinus2(index int) int {
	if baseIndex := index - 2; baseIndex > 0 {
		return baseIndex
	} else {
		return 0
	}
}

func sumSweeps(sweeps []int) int {
	return sweeps[0] +sweeps[1] + sweeps[2]
}

func main() {

	var sweeps [][]int
	sweepIncreases := 0
	scanner := internal.GetScannerForInput(internal.FILE()+"input")
	scanner.Scan()
	sweeps = append(sweeps, []int{})
	sweeps[0] = append(sweeps[0], internal.StringToInt(scanner.Text()))
	currentSweepIndex := 1
	for scanner.Scan() {
		nextSweepValue := internal.StringToInt(scanner.Text())
		sweeps = append(sweeps, []int{})
		sweeps[currentSweepIndex] = append(sweeps[currentSweepIndex], nextSweepValue)
		sweepIndex := zeroOrIMinus2(currentSweepIndex)
		for i := currentSweepIndex-1; i>=sweepIndex;i-- {
			sweeps[i] = append(sweeps[i], nextSweepValue)
		}
		if len(sweeps) > 3 {
			if len(sweeps[currentSweepIndex-3]) == len(sweeps[currentSweepIndex-2]) {
				if sumSweeps(sweeps[currentSweepIndex-3]) < sumSweeps(sweeps[currentSweepIndex-2]) {
					sweepIncreases++
				}
			}
		}
		currentSweepIndex+=1
	}
	fmt.Println(sweepIncreases)
}
