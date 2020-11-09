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

	var totalTestCases, answer int
	fmt.Fscanf(in, "%d\n", &totalTestCases)
	arrayOfAltitude := make([]int, totalTestCases+1)

	for i := 0; i < totalTestCases; i++ {
		var roadType rune
		var roadTypeToInteger int
		fmt.Fscanf(in, "%c ", &roadType)

		roadTypeToInteger = convertRoadTypeToInteger(roadType)

		if i == 0 {
			arrayOfAltitude[i] = roadTypeToInteger
		} else {
			arrayOfAltitude[i] = arrayOfAltitude[i-1] + roadTypeToInteger	

			if isInZeroAltitude(arrayOfAltitude[i]) {
				if isFromDownhill(arrayOfAltitude[i-1]) {
					answer++
				}
			}
		}
	}

	fmt.Fprintln(out, answer)
}

func convertRoadTypeToInteger(roadType rune) int {
	if roadType == 'U' {
		return 1
	}

	return -1
}

func isInZeroAltitude(altitude int) bool {
	if altitude == 0 {
		return true
	}

	return false
}

func isFromDownhill(altitude int) bool {
	if altitude < 0 {
		return true
	}

	return false
}