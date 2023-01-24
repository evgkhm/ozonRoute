package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var teams int //количество наборов входных данных (3)

	fmt.Fscan(in, &teams)

	for i := 0; i < teams; i++ {
		developsCount := 0 //кол-во разработчиков (6)
		fmt.Fscan(in, &developsCount)
		var developsSlice []int
		//developsSlice := make([]int, 0)
		for j := 1; j <= developsCount; j++ {
			develop := 0 //сами разработчики 2 1 3 1 1 4
			fmt.Fscan(in, &develop)
			developsSlice = append(developsSlice, develop)
		}
		var developsNumbers []int
		for j := 0; j < developsCount; j++ {
			developsNumbers = append(developsNumbers, j+1)
		}
		//developsNumbers := make([]int, len(developsSlice))
		/*for j := range developsNumbers {
			developsNumbers[j] = j + 1
		}*/

		for len(developsSlice) > 0 {
			resultSlice := []int{developsNumbers[0]}
			min := 100
			minIndex := 0

			for j := 1; j < len(developsSlice); j++ {
				diff := int(math.Abs(float64(developsSlice[0] - developsSlice[j])))
				if diff < min {
					min = diff
					minIndex = j
				}
			}
			resultSlice = append(resultSlice, developsNumbers[minIndex])
			fmt.Fprintln(out, resultSlice[0], resultSlice[1])
			developsSlice = append(developsSlice[:minIndex], developsSlice[minIndex+1:]...)
			developsSlice = developsSlice[1:]
			developsNumbers = append(developsNumbers[:minIndex], developsNumbers[minIndex+1:]...)
			developsNumbers = developsNumbers[1:]
		}
		fmt.Fprintln(out)
	}
}
