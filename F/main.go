package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var otrezki int
	fmt.Fscan(in, &otrezki) //6

	for i := 0; i < otrezki; i++ {
		var otrezok int //5
		fmt.Fscan(in, &otrezok)
		res := "YES"
		for j := 0; j < otrezok; j++ {
			var date, rawLeftDate, rawRightDate string
			fmt.Fscan(in, &date)

			tmp := strings.Split(date, "-")
			rawLeftDate = tmp[0]
			rawRightDate = tmp[1]

			layout := "15:04:05"
			leftTime, _ := time.Parse(layout, rawLeftDate)
			rightTime, _ := time.Parse(layout, rawRightDate)

			if leftTime.Before(rightTime) == true || leftTime.Equal(rightTime) == true {
				//res = "YES"
			} else {
				res = "NO"
				continue
			}
			//rightSecs := t.Second()

			//intDate, _ := strconv.Atoi(date)
			var timeSlice []time.Time
			if otrezok > 1 {
				timeSlice = append(timeSlice, leftTime)
				timeSlice = append(timeSlice, rightTime)

				sort.Slice(timeSlice, func(i, j int) bool {
					return timeSlice[i].Before(timeSlice[j])
				})
			}
			fmt.Fprintln(out, timeSlice)
		}
		fmt.Fprintln(out, res)
	}

}
