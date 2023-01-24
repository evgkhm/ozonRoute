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

	var tableCount int
	fmt.Fscan(in, &tableCount)

	for i := 0; i < tableCount; i++ {
		var slice [][]int
		var stringTable, columnTable int

		fmt.Fscan(in, &stringTable, &columnTable)
		for n := 0; n < stringTable; n++ {
			var tmpSlice []int
			for k := 0; k < columnTable; k++ {
				value := 0
				fmt.Fscan(in, &value)
				tmpSlice = append(tmpSlice, value)
			}
			slice = append(slice, tmpSlice)
		}

		var clicks int
		fmt.Fscan(in, &clicks)
		for clicks > 0 {
			var sortStolbec int //какой столбец нужно сортировать
			fmt.Fscan(in, &sortStolbec)
			j := 0
			for j = 0; j < stringTable-1; j++ {
				if slice[j][sortStolbec-1] > slice[j+1][sortStolbec-1] {
					//slice[i][sortStolbec], slice[i][sortStolbec+1] = slice[i][sortStolbec+1],slice[i][sortStolbec]
					slice[j], slice[j+1] = slice[j+1], slice[j]
					j = 0
				}
			}
			clicks--
		}
		for _, stroka := range slice {
			for _, val := range stroka {
				fmt.Fprint(out, val, " ")
			}
			fmt.Fprintln(out, " ")
		}
	}

}
