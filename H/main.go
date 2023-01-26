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

	//res := "YES"
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

		res := Calc(slice)

		fmt.Fprintln(out, res)
	}
}

func FindAndDeleteRight(cellN int, cellM int, slice [][]string, targetCell string) [][]string {
	if cellM >= stolbec || cellN >= stroka {
		return slice
	}

	currCell := slice[cellN][cellM]

	if currCell == "." {
		cellM++
		FindAndDeleteRight(cellN, cellM, slice, targetCell)
	}
	//переход на след строку
	if cellM >= stolbec || currCell != targetCell {
		cellM = 0
		cellN += 1
		FindAndDeleteRight(cellN, cellM, slice, targetCell)
	}

	if cellM >= stolbec || cellN >= stroka {
		return slice
	}
	slice[cellN][cellM] = "."

	//переход на след столбец
	cellM += 2

	FindAndDeleteRight(cellN, cellM, slice, targetCell)
	return slice
}

func Calc(slice [][]string) string {
	res := "YES"
	targetCell := slice[0][0]
	cellN := 0
	cellM := 0
	slice = FindAndDeleteRight(cellN, cellM, slice, targetCell)
	for _, str := range slice {
		for _, val := range str {
			if val == targetCell {
				res = "NO"
			}
		}
	}

	return res
}
