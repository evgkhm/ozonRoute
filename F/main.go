package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Interval []struct {
	leftTime, rightTime int
}

func (i Interval) Len() int {
	return len(i)
}

func (i Interval) Less(k, j int) bool {
	//return i[k].leftTime.Before(i[j].leftTime)
	return i[k].leftTime < i[j].leftTime
}

func (i Interval) Swap(k, j int) {
	i[k], i[j] = i[j], i[k]
}

func Check(rawDate string) bool {
	raw := strings.Split(rawDate, ":")
	sec, _ := strconv.Atoi(raw[2])
	min, _ := strconv.Atoi(raw[1])
	hour, _ := strconv.Atoi(raw[0])

	if sec > 59 || min > 59 || hour > 23 {
		return false
	}
	return true
}

func AnotherCheck(leftTime, rightTime time.Time) bool {
	if leftTime.Before(rightTime) == true || leftTime.Equal(rightTime) == true {
		//res = "YES"
	} else {
		//res = "NO"
		//continue
		return false
	}
	return true
}

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

		var intervals = make(Interval, otrezok)
		for j := 0; j < otrezok; j++ {
			var date, rawLeftDate, rawRightDate string
			fmt.Fscan(in, &date)

			tmp := strings.Split(date, "-")
			rawLeftDate = tmp[0]
			rawRightDate = tmp[1]
			check := Check(rawLeftDate)
			if check == false {
				res = "NO"
				continue
			}
			check = Check(rawRightDate)
			if check == false {
				res = "NO"
				continue
			}

			layout := "15:04:05"
			leftTime, _ := time.Parse(layout, rawLeftDate)
			rightTime, _ := time.Parse(layout, rawRightDate)

			check = AnotherCheck(leftTime, rightTime)
			if check == false {
				res = "NO"
				continue
			}
			//перевод в секунды
			intervals[j].leftTime = leftTime.Second() + (leftTime.Minute() * 60) + (leftTime.Hour() * 60 * 60)
			intervals[j].rightTime = rightTime.Second() + (rightTime.Minute() * 60) + (rightTime.Hour() * 60 * 60)
		}

		if otrezok > 1 {
			sort.Sort(intervals)
			for n := 0; n < len(intervals)-1; n++ {
				if intervals[n].rightTime >= intervals[n+1].leftTime {
					res = "NO"
					continue
				}
			}
		}
		fmt.Fprintln(out, res)
	}
}
