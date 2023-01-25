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

	for i := 0; i < employee; i++ {
		result := "YES"
		var days int //5
		fmt.Fscan(in, &days)

		prevReport := 0
		m := make(map[int]bool) //мапа с отчетами

		for j := 0; j < days; j++ {
			var report int
			fmt.Fscan(in, &report)

			_, ok := m[report]
			if ok && prevReport == report {
				continue
			} else if ok && prevReport != report {
				result = "NO"
				continue
			} else {
				m[report] = true
			}
			prevReport = report
		}
		fmt.Fprintln(out, result)
	}

}
