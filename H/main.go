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
	nextTargetCell             string
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

func FindAndDelete(cells *Coordinates, slice [][]string) [][]string {
	processed := make(map[Coordinates]bool) //обработанные ячейки
	//queue := list.New()
	//queue.PushBack()
	for n, value := range slice {
		for m, val := range value {
			cells.cellN = n
			cells.cellM = m
			if _, ok := processed[*cells]; ok {
				continue
			}
			processed[*cells] = true

			if val != cells.targetCell && val != "." && cells.foundNext == false { //конец обхода с записей нового цвета
				cells.nextTargetCell = val
				cells.foundNext = true
				break //выход из этой строки
			} else if val != cells.targetCell && val != "." { //конец обхода
				break //выход из этой строки
			}

			if val == cells.targetCell {
				slice[n][m] = "."
				//FindAndDelete(cells, slice, processed)
			}
		}
	}

	return slice
}

func Calc(cells *Coordinates, slice [][]string) string {
	res := "YES"

	slice = FindAndDelete(cells, slice)
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
	//cells.cellN = cells.nextN
	//cells.cellM = cells.nextM
	cells.targetCell = cells.nextTargetCell
	cells.foundNext = false
	return res
}
