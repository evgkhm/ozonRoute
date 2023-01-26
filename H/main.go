package main

import (
	"bufio"
	"fmt"
	"os"
)

var stroka, stolbec int

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	res := "YES"
	for i := 0; i < testCount; i++ {

		fmt.Fscan(in, &stroka, &stolbec)

		slice := make([][]string, stroka)
		for n := range slice {
			slice[n] = make([]string, stolbec)
		}

		//append input string into 1 dim slice
		var tmpSlice []string
		for j := 0; j < stroka; j++ {
			var str string
			fmt.Fscan(in, &str)
			tmpSlice = append(tmpSlice, str)
		}

		//from 1 dim slice to 2 dim
		n := 0
		for _, str := range tmpSlice {
			m := 0
			for _, val := range str {
				slice[n][m] = string(val)
				m++
			}
			n++
		}

		res = Calc(slice)

		fmt.Fprintln(out, slice)
	}
}

func FindAndDelete(cellN int, cellM int, slice [][]string) (int, int, [][]string) {
	/*if currCell != prevCell {
		break
	}*/
	currCell := slice[cellN][cellM]
	return FindAndDelete(cellN+2, cellM+2, slice)
}

func Calc(slice [][]string) string {
	res := "YES"
	prevSymbol := slice[0][0]
	cellN := 0
	cellM := 0
	FindAndDelete(cellN, cellM, slice)
	/*
		var region []
		for n := 0; n < stroka; n++ {
			for m := 2; m < stolbec; m++ {
				if slice[n][m] == prevSymbol {

				}
			}
		}*/
	return res
}
