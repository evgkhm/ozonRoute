package main

import (
	"bufio"
	"fmt"
	"os"
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
		//res := "YES"
		var otrezok int //5
		fmt.Fscan(in, &otrezok)

		//prevReport := 0
		//m := make(map[int]bool) //мапа с отчетами
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
			}
			//rightSecs := t.Second()
			//fmt.Fprintln(out, diff)
			//fmt.Fprintln(out, leftDate, rightDate)

		}
		fmt.Fprintln(out, res)
	}

}
