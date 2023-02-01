package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type CPU struct {
	Energy    int
	StartTime int
	TimeExec  int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var CPUs, tasks int
	fmt.Fscan(in, &CPUs, &tasks)

	var cpuSlice []*CPU

	for i := 0; i < CPUs; i++ {
		var energy int
		fmt.Fscan(in, &energy)
		cpuSlice = append(cpuSlice, &CPU{Energy: energy})
	}

	sort.Slice(cpuSlice, func(i, j int) bool {
		return cpuSlice[i].Energy < cpuSlice[j].Energy
	})

	res := 0
	for i := 0; i < tasks; i++ {
		var startTime, timeDuration int
		fmt.Fscan(in, &startTime, &timeDuration)
		for _, val := range cpuSlice {
			if startTime-val.StartTime >= val.TimeExec {
				res += val.Energy * timeDuration
				val.StartTime = startTime
				val.TimeExec = timeDuration
				break
			}
		}

	}
	fmt.Fprintln(out, res)
}
