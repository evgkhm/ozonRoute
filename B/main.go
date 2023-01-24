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

	var productCount int //количество наборов входных данных (6)

	fmt.Fscan(in, &productCount)

	for i := 0; i < productCount; i++ {
		var inputData int      //наборы входных данных (12 ... 12 ... 1...)
		var product int        //сами продукты (2 2 2 2 2 2 2 3 3 3 3 3)
		m := make(map[int]int) //мапа с повторяющимися эл-тами

		fmt.Fscan(in, &inputData)

		for j := 0; j < inputData; j++ {
			fmt.Fscan(in, &product)
			_, ok := m[product]
			if ok {
				m[product]++
			} else {
				m[product] = 1
			}

		}
		res := 0
		for key, value := range m {
			res += (value - (value / 3)) * key
		}
		fmt.Fprintln(out, res)
	}

}
