package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Person struct {
	ID      int
	Friends Friends
}

type Friends struct {
	ID    []int
	Count []int
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

	for i := 1; i <= friends; i++ {
		commonFriends := make(map[int]int)

		if _, ok := m[i]; ok { //текущий чел
			for _, val := range m[i] { //ind-ключ(сам чел) val - его друг
				if _, okk := m[val]; okk { //его друг
					for _, friendOfFriend := range m[val] { //значения его друга
						if friendOfFriend == i { //совпадение с начальным другом (1)
							continue
						} else {
							//добавление в мапу возможного друга и увеличить счетчик
							if _, okey := commonFriends[friendOfFriend]; !okey {
								commonFriends[friendOfFriend] = 1
							} else {
								counter := commonFriends[friendOfFriend]
								counter++
								commonFriends[friendOfFriend] = counter
							}
						}
					}
				}
			}

			//удаление повторяющихся друзей
			for _, val := range m[i] {
				for key, _ := range commonFriends {
					if val == key {
						delete(commonFriends, key)
					}
				}
			}

			//нет общих друзей
			if len(commonFriends) == 0 {
				fmt.Fprintln(out, 0)
				continue
			}

			maxVal := 0
			for _, val := range commonFriends {
				if val > maxVal {
					maxVal = val
				}
			}

			var sliceForPrint []int
			for key, val := range commonFriends {
				if val == maxVal {
					//fmt.Fprintln(out, key)
					sliceForPrint = append(sliceForPrint, key)
				}
			}
			sort.Slice(sliceForPrint, func(i, j int) bool {
				return sliceForPrint[i] < sliceForPrint[j]
			})
			for _, v := range sliceForPrint {
				fmt.Fprint(out, v, " ")
			}
			fmt.Fprintln(out)
		} else {
			fmt.Fprintln(out, 0)
		}
	}

}
