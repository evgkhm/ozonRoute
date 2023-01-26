package main

import (
	"bufio"
	"fmt"
	"os"
)

var stroka, stolbec int

type Coordinates struct {
	cellN, cellM, nextN, nextM int
	targetCell                 string
}

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

		//start := slice[0][0]
		cells := &Coordinates{}
		cells.targetCell = slice[0][0]

		endStr := stroka - 1
		endStolb := stolbec - 1
		res := "YES"
		for slice[endStr][endStolb] != "." && res == "YES" {
			res = Calc(cells, slice)
		}

		fmt.Fprintln(out, res)
	}
}

func FindAndDeleteRight(cells *Coordinates, slice [][]string) [][]string {
	if cells.cellM >= stolbec || cells.cellN >= stroka {
		return slice
	}

	currCell := slice[cells.cellN][cells.cellM]

	if currCell == "." {
		cells.cellM++
		FindAndDeleteRight(cells, slice)
	} else if currCell != cells.targetCell { //попадание в след символ
		cells.nextN = cells.cellN
		cells.nextM = cells.cellM
	}
	//переход на след строку
	if cells.cellM >= stolbec || currCell != cells.targetCell {
		cells.cellM = 0
		cells.cellN += 1
		FindAndDeleteRight(cells, slice)
	}

	if cells.cellM >= stolbec || cells.cellN >= stroka {
		return slice
	}
	slice[cells.cellN][cells.cellM] = "."

	//переход на след столбец
	cells.cellM += 2

	FindAndDeleteRight(cells, slice)
	return slice
}

func Calc(cells *Coordinates, slice [][]string) string {
	res := "YES"
	//targetCell := slice[0][0]
	//cellN := 0
	//cellM := 0
	//cells := &Coordinates{}
	slice = FindAndDeleteRight(cells, slice)
	//проверка что нет островов
	for _, str := range slice {
		for _, val := range str {
			if val == cells.targetCell {
				res = "NO"
				return res
			}
		}
	}
	cells.cellN = cells.nextN
	cells.cellM = cells.nextM
	cells.targetCell = slice[cells.nextN][cells.nextM]
	return res
}
