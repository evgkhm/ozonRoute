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

		cells := &Coordinates{}
		cells.targetCell = slice[0][0]

		res := "YES"
		var pointer = &res

		Calc(cells, slice, pointer)

		fmt.Fprintln(out, res)
	}
}

func FindAndDelete(cells *Coordinates, slice [][]string, processed map[Coordinates]bool, res string) ([][]string, string) {
L1:
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
						FindAndDelete(cells, slice, processed, res)
					}
				}
				if m+2 < stolbec { //сосед справа
					if slice[n][m+2] == cells.targetCell {
						cells.cellN = n
						cells.cellM = m + 2
						FindAndDelete(cells, slice, processed, res)
					}
				}
				if m+1 < stolbec && n+1 < stroka { //сосед справа внизу
					if slice[n+1][m+1] == cells.targetCell {
						cells.cellN = n + 1
						cells.cellM = m + 1
						FindAndDelete(cells, slice, processed, res)
					}
				}
				if m-1 >= 0 && n+1 < stroka { //сосед слева внизу
					if slice[n+1][m-1] == cells.targetCell {
						cells.cellN = n + 1
						cells.cellM = m - 1
						FindAndDelete(cells, slice, processed, res)
					}
				}
				if m-2 >= 0 { //сосед слева
					if slice[n][m-2] == cells.targetCell {
						cells.cellN = n
						cells.cellM = m - 2
						FindAndDelete(cells, slice, processed, res)
					}
				}

				if m-1 >= 0 && n-1 >= 0 { //сосед слева наверху
					if slice[n-1][m-1] == cells.targetCell {
						cells.cellN = n - 1
						cells.cellM = m - 1
						FindAndDelete(cells, slice, processed, res)
					}
				}
			}
			break L1 //нужно полностью выйти из цикла
		}
	}

	return slice, res
}

func Calc(cells *Coordinates, slice [][]string, p *string) *string {
	processed := make(map[Coordinates]bool) //обработанные ячейки

	FindAndDelete(cells, slice, processed, *p)

	//проверка что нет островов
	for _, str := range slice {
		for _, val := range str {
			if val == cells.targetCell {
				*p = "NO"
				return p
			}
		}
	}

	//запуск заново если есть символы кроме "."
	for n, str := range slice {
		for m, val := range str {
			if val != "." {
				cells.cellN = n
				cells.cellM = m
				cells.targetCell = val
				Calc(cells, slice, p)
			}
		}
	}

	return p
}
