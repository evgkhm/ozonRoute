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

	for n, value := range slice {
		for m, val := range value {
			cells.cellN = n
			cells.cellM = m

			if _, ok := processed[*cells]; ok {
				continue
			}
			processed[*cells] = true

			if val != cells.targetCell && val != "." && cells.foundNext == false { //новый цвет
				cells.nextTargetCell = val
				cells.foundNext = true
				continue
			}

			if val == "." {
				continue
			}
			if val == cells.targetCell {
				isNeighbors := false
				slice[n][m] = "."

				if m+1 < stolbec && n-1 > 0 { //сосед справа наверху
					cells.cellN = n - 1
					cells.cellM = m + 1
					processed[*cells] = true
					if slice[n-1][m+1] == cells.targetCell {
						slice[n-1][m+1] = "."
						isNeighbors = true
					}
				}
				if m+2 < stolbec { //сосед справа
					cells.cellM = m + 2
					processed[*cells] = true
					if slice[n][m+2] == cells.targetCell {
						slice[n][m+2] = "."
						isNeighbors = true
					}
				}
				if m+1 < stolbec && n+1 < stroka { //сосед справа внизу
					cells.cellN = n + 1
					cells.cellM = m + 1
					processed[*cells] = true
					if slice[n+1][m+1] == cells.targetCell {
						slice[n+1][m+1] = "."
						isNeighbors = true
					}
				}
				if n+2 < stroka { //сосед внизу
					cells.cellN = n + 2
					processed[*cells] = true
					if slice[n+2][m] == cells.targetCell {
						slice[n+2][m] = "."
						isNeighbors = true
					}
				}
				if m-1 > 0 && n+1 < stroka { //сосед слева внизу
					cells.cellN = n + 1
					cells.cellM = m - 1
					processed[*cells] = true
					if slice[n+1][m-1] == cells.targetCell {
						slice[n+1][m-1] = "."
						isNeighbors = true

					}
				}
				if m-2 > 0 { //сосед слева
					cells.cellM = m - 2
					processed[*cells] = true
					if slice[n][m-2] == cells.targetCell {
						slice[n][m-2] = "."
						isNeighbors = true
					}
				}

				if m-1 > 0 && n-1 > 0 { //сосед слева наверху
					cells.cellN = n - 1
					cells.cellM = m - 1
					processed[*cells] = true
					if slice[n-1][m-1] == cells.targetCell {
						slice[n-1][m-1] = "."
						isNeighbors = true
					}
				}

				if n-2 > 0 { //сосед наверху
					cells.cellN = n - 2
					processed[*cells] = true
					if slice[n-2][m] == cells.targetCell {
						slice[n-2][m] = "."
						isNeighbors = true
					}
				}

				if isNeighbors == false { //нет соседей, выход из функции
					return slice
				}
				continue
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
