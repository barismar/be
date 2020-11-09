package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	var mapCountTotalRepetition map[int]int
	var totalTestCases, answer int
	
	mapCountTotalRepetition = make(map[int]int)
	fmt.Fscanf(in, "%d\n", &totalTestCases)

	for i := 0; i < totalTestCases; i++ {
		var sockColour int
		fmt.Fscanf(in, "%d ", &sockColour)

		mapCountTotalRepetition[sockColour]++

		if isSocksMatched(mapCountTotalRepetition[sockColour]) {
			answer++
		}
	}

	fmt.Fprintln(out, answer)
}

func isSocksMatched(totalSocks int) bool {
	if totalSocks % 2 == 0 {
		return true
	}

	return false
}
