package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// Item - это то, чем мы управляем в приоритетной очереди.
type Item struct {
	value    string // Значение элемента; произвольное.
	priority int    // Приоритет элемента в очереди.
	// Индекс необходим для обновления
	// и поддерживается методами heap.Interface.
	index int // Индекс элемента в куче.
}

// PriorityQueue реализует heap.Interface и содержит Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// Мы хотим, чтобы Pop давал нам самый высокий,
	// а не самый низкий приоритет,
	// поэтому здесь мы используем оператор больше.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // избежать утечки памяти
	item.index = -1 // для безопасности
	*pq = old[0 : n-1]
	return item
}

// update изменяет приоритет и значение Item в очереди.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var CPUs, tasks int
	fmt.Fscan(in, &CPUs, &tasks)

	var energyConsump []int

	for i := 0; i < CPUs; i++ {
		var energy int
		fmt.Fscan(in, &energy)
		energyConsump = append(energyConsump, energy)
	}

	for i := 0; i < tasks; i++ {
		var task, sec int
		fmt.Fscan(in, &task, &sec)

		//fmt.Fprintln(out, n+m)
	}
}
