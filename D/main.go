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
		stringTable -= 1 //начало с 0 индекса
		for clicks > 0 {
			var sortStolbec int //какой столбец нужно сортировать
			fmt.Fscan(in, &sortStolbec)
			sortStolbec -= 1

			j := 0
			for j < stringTable {
				if slice[j][sortStolbec] > slice[j+1][sortStolbec] {
					slice[j], slice[j+1] = slice[j+1], slice[j]
					j = 0 //???
				} else {
					j++
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
		fmt.Fprintln(out, " ")
	}
}
