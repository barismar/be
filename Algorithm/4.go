package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n := 100
	arr := make([]int, n + 1)
	var answer int

	for i := 1; i <= n; i++ {
		for j := i; j <= n; j += i {
			if arr[j] == 0 {
				arr[j] = 1
				answer++
			} else {
				arr[j] = 0
				answer--
			}
		}
	}
	
	fmt.Fprintln(out, answer)
}