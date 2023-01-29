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

		//endStr := stroka - 1
		//endStolb := stolbec - 1
		res := "YES"
		res = Calc(cells, slice)
		/*for slice[endStr][endStolb] != "." && res == "YES" {

		}*/

		fmt.Fprintln(out, res)
	}
}

func FindAndDelete(cells *Coordinates, slice [][]string, processed map[Coordinates]bool) [][]string {

	for n := cells.cellN; n < stroka; n++ {
		for m := cells.cellM; m < stolbec; m++ {
			cells.cellN = n
			cells.cellM = m
			if _, ok := processed[*cells]; ok {
				continue
			}
			processed[*cells] = true

			val := slice[n][m]

			if val == "." {
				continue
			}

			if val == cells.targetCell {
				slice[n][m] = "."

				if m+1 < stolbec && n-1 >= 0 { //сосед справа наверху
					if slice[n-1][m+1] == cells.targetCell {
						cells.cellN = n - 1
						cells.cellM = m + 1
						FindAndDelete(cells, slice, processed)
					}
				}
				if m+2 < stolbec { //сосед справа
					if slice[n][m+2] == cells.targetCell {
						cells.cellN = n
						cells.cellM = m + 2
						FindAndDelete(cells, slice, processed)
					}
				}
				if m+1 < stolbec && n+1 < stroka { //сосед справа внизу
					if slice[n+1][m+1] == cells.targetCell {
						cells.cellN = n + 1
						cells.cellM = m + 1
						FindAndDelete(cells, slice, processed)
					}
				}
				if n+2 < stroka { //сосед внизу
					if slice[n+2][m] == cells.targetCell {
						cells.cellN = n + 2
						cells.cellM = m
						FindAndDelete(cells, slice, processed)
					}
				}
				if m-1 >= 0 && n+1 < stroka { //сосед слева внизу
					if slice[n+1][m-1] == cells.targetCell {
						cells.cellN = n + 1
						cells.cellM = m - 1
						FindAndDelete(cells, slice, processed)
					}
				}
				if m-2 >= 0 { //сосед слева
					if slice[n][m-2] == cells.targetCell {
						cells.cellN = n
						cells.cellM = m - 2
						FindAndDelete(cells, slice, processed)
					}
				}

				if m-1 >= 0 && n-1 >= 0 { //сосед слева наверху
					if slice[n-1][m-1] == cells.targetCell {
						cells.cellN = n - 1
						cells.cellM = m - 1
						FindAndDelete(cells, slice, processed)
					}
				}

				if n-2 >= 0 { //сосед наверху
					if slice[n-2][m] == cells.targetCell {
						cells.cellN = n - 2
						cells.cellM = m
						FindAndDelete(cells, slice, processed)
					}
				}
				return slice
			}
		}
	}

	return slice
}

func Calc(cells *Coordinates, slice [][]string) string {
	res := "YES"

	processed := make(map[Coordinates]bool) //обработанные ячейки

L1:
	for n, str := range slice {
		for m, val := range str {
			if val != "." {
				cells.cellN = n
				cells.cellM = m
				cells.targetCell = val
				break L1
			}
		}
	}
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

	return res
}
