package main

import (
	"bufio"
	"fmt"
	"os"
)

func validateMap(mapData [][]rune) bool {
	// Create a 2D boolean array to mark visited hexagons
	visited := make([][]bool, len(mapData))
	for i := range visited {
		visited[i] = make([]bool, len(mapData[i]))
	}

	// Create a map to store the regions on the map
	regions := make(map[rune]bool)

	// Iterate through the mapData array to find all regions on the map
	for i := range mapData {
		for j := range mapData[i] {
			if mapData[i][j] != '.' && !regions[mapData[i][j]] {
				regions[mapData[i][j]] = true
			}
		}
	}

	// Define a helper function to check if a hexagon is valid
	var dfs func(int, int, rune) bool
	dfs = func(i, j int, region rune) bool {
		if i < 0 || i >= len(mapData) || j < 0 || j >= len(mapData[i]) || visited[i][j] || mapData[i][j] != region {
			return true
		}

		visited[i][j] = true
		valid := dfs(i+1, j, region) && dfs(i-1, j, region) && dfs(i, j+1, region) && dfs(i, j-1, region) &&
			dfs(i+1, j+1, region) && dfs(i-1, j-1, region)
		visited[i][j] = false
		return valid
	}

	// Iterate through all regions on the map and check if they are connected
	for region := range regions {
		regionConnected := false
		for i := range mapData {
			for j := range mapData[i] {
				if !visited[i][j] && mapData[i][j] == region {
					if !regionConnected {
						regionConnected = true
					} else if !dfs(i, j, region) {
						return false
					}
				}
			}
		}
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscan(in, &testCount)

	//res := "YES"
	for i := 0; i < testCount; i++ {
		var stroka, stolbec int
		fmt.Fscan(in, &stroka, &stolbec)

		slice := make([][]rune, stroka)
		for n := range slice {
			slice[n] = make([]rune, stolbec)
		}

		//append input string into 1 dim slice
		var tmpSlice []rune
		for j := 0; j < stroka; j++ {
			var str rune
			fmt.Fscan(in, &str)
			tmpSlice = append(tmpSlice, str)
		}

		//from 1 dim slice to 2 dim
		n := 0
		for _, str := range tmpSlice {
			m := 0
			for _, val := range string(str) {
				slice[n][m] = val
				m++
			}
			n++
		}

		//start := slice[0][0]
		//cells := &Coordinates{}
		//cells.targetCell = slice[0][0]

		//endStr := stroka - 1
		//endStolb := stolbec - 1

		//for slice[endStr][endStolb] != "." && res == "YES" {
		//	res = Calc(cells, slice)
		//}
		if validateMap(slice) == true {
			res := "YES"
			fmt.Fprintln(out, res)
		} else {
			res := "NO"
			fmt.Fprintln(out, res)
		}

	}
}
