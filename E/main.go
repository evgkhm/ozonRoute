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

	var employee int
	fmt.Fscan(in, &employee) //5
	result := "N0"

	for i := 0; i < employee; i++ {
		var days int //5
		fmt.Fscan(in, &days)

		prevReport := 0
		m := make(map[int]bool) //мапа с отчетами
	L1:
		for j := 0; j < days; j++ {
			var report int
			fmt.Fscan(in, &report)

			_, ok := m[report]
			if ok && prevReport == report {
				continue
			} else if ok && prevReport != report {
				result = "N0"
				break L1
			} else {
				m[report] = true
			}
			prevReport = report

			result = "YES"
		}
		fmt.Fprintln(out, result)
	}

}
