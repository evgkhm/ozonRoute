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
	foundNext                  bool
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

func FindAndDelete(cells *Coordinates, slice [][]string, processed map[int]bool) [][]string {
	count := 0
	for n, value := range slice {
		for m, val := range value {
			//if processed[n][m] {
			//	continue // не перебираем одну вершину два раза
			//}
			if _, ok := processed[m]; ok {
				continue
			}
			processed[m] = true

			if val == cells.targetCell { //удаление из слайса
				slice[n][m] = "."
				FindAndDelete(cells, slice, processed)
			}

			if val != cells.targetCell && val != "." { //конец обхода
				cells.nextN = n
				cells.nextN = m
				return slice
			}
		}
	}

	return slice
}

func Calc(cells *Coordinates, slice [][]string) string {
	res := "YES"
	processed := make(map[int]bool) //обработанные ячейки

	slice = FindAndDelete(cells, slice, processed)
	//проверка что нет островов
	for _, str := range slice {
		for _, val := range str {
			if val == cells.targetCell {
				res = "NO"
				return res
			}
		}
	}
	//новый цикл поиска
	cells.cellN = cells.nextN
	cells.cellM = cells.nextM
	cells.targetCell = slice[cells.nextN][cells.nextM]
	cells.foundNext = false
	return res
}
