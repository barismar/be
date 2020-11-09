package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscanf(in, "%d\n", &n)
	t := strconv.Itoa(n)
	stringLength := len(t) - 1

	for i := 0; i < len(t); i++ {
		answer := string(t[i])
		for j := 0; j < stringLength; j++ {
			answer += "0"
		}
		stringLength--
		fmt.Fprintln(out, answer)
	}
}