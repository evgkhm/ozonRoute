package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	ID int
	IDFriend int
	Count int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var friends int
	fmt.Fscan(in, &friends)

	var couples int
	fmt.Fscan(in, &couples)

	m := make(map[int][]int) //мапа со слайсами
	for i := 0; i < couples; i++ {
		var first, second int
		fmt.Fscan(in, &first, &second)

		m[first] = append(m[first], second)
		m[second] = append(m[second], first)
	}

	for _, val := range m { //сортировка в слайсах
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
	}

	for i := 1; i < couples; i++ {
		var person Person{}

		if _, ok := m[i]; ok { //текущий чел
			for _, val := range m[i] { //val - его друг

				if _, okk := m[val]; okk { //его друг
					for _, friendOfFriend := range m[val] { //значения его друга
						if friendOfFriend == i { //совпадение с начальным другом (1)
							continue
						} else {
							//TODO: дбавить в слайс счетчик
						}
					}


				} else {

				}
			}
		} else {
			fmt.Println(out, 0)
		}
	}

	fmt.Println(out, m)
}
