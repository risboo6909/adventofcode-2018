package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	next   []*Node
	id     string
	refCnt uint
	index  int // the index of the item in the heap
}

func (node Node) String() string {
	return node.id
}

type Worker struct {
	timeLeft uint
	task     *Node
}

func (worker Worker) String() string {
	return fmt.Sprintf("%d", worker.timeLeft)
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].id <= pq[j].id
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Node, id string) {
	item.id = id
	heap.Fix(pq, item.index)
}

func getStdin() []string {
	var output []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, line)
	}

	return output
}

func parseInput(lines []string) map[string]*Node {

	refs := make(map[string]*Node)

	for _, line := range lines {

		parts := strings.Split(line, " ")

		// this task must be finished before...
		prevTaskId := parts[1]
		// ...this task starts
		taskId := parts[7]

		connectNodes(taskId, prevTaskId, refs)
	}

	return refs

}

func connectNodes(taskId string, prevTaskId string, refs map[string]*Node) {

	var prevNode *Node

	// if prev node exists get the reference to it, create a new one otherwise
	if node, ok := refs[prevTaskId]; ok {
		prevNode = node
	} else {
		prevNode = &Node{next: nil, id: prevTaskId}
		refs[prevTaskId] = prevNode
	}

	// if cur node exists get the reference to it, create a new one otherwise
	// connect previous node to the current node
	if node, ok := refs[taskId]; ok {
		prevNode.next = append(prevNode.next, node)
		node.refCnt++
	} else {
		node := Node{next: nil, id: taskId}
		prevNode.next = append(prevNode.next, &node)
		refs[taskId] = &node
		node.refCnt++
	}
}

func findOrigins(refs map[string]*Node) []*Node {

	// root node - is the one that doesn't depend on anyone else
	// i.e. its refCnt == 0

	var roots []*Node

	for _, v := range refs {
		if v.refCnt == 0 {
			roots = append(roots, v)
		}
	}

	return roots
}

func serialExecution(roots []*Node) string {

	var result []string

	// init priority queue
	pq := make(PriorityQueue, len(roots))
	for idx := range roots {
		pq[idx] = roots[idx]
	}

	heap.Init(&pq)

	// traverse all jobs
	for pq.Len() > 0 {
		curJob := heap.Pop(&pq).(*Node)

		for idx := range curJob.next {
			nextJob := curJob.next[idx]
			nextJob.refCnt--
			if nextJob.refCnt == 0 {
				heap.Push(&pq, curJob.next[idx])
			}
		}
		result = append(result, curJob.id)
	}

	return strings.Join(result, "")
}

func freeWorkerIdx(workers []*Worker) int {
	for idx, worker := range workers {
		if worker.timeLeft == 0 {
			return idx
		}
	}
	return -1
}

func updateWorkers(workers []*Worker) {
	for idx := range workers {
		if workers[idx].timeLeft > 0 {
			workers[idx].timeLeft--
		}
	}
}

func getEstimatedTime(curJob *Node) uint {
	return 60 + (uint(curJob.id[0]) - 64)
}

func parallelExecution(roots []*Node, workers_no int) int {

	last_finished_iter := 0
	iter := 0

	var result []string

	var workers []*Worker
	// init workers
	for x := 0; x < workers_no; x++ {
		workers = append(workers, &Worker{timeLeft: 0})
	}

	// init priority queue
	pq := make(PriorityQueue, len(roots))
	for idx := range roots {
		pq[idx] = roots[idx]
	}

	// traverse all jobs
	for (iter - last_finished_iter) < 500 {

		freeWorkerIdx := freeWorkerIdx(workers)

		// assign new job only if we have jobs pending and a free worker
		if pq.Len() > 0 && freeWorkerIdx != -1 {
			curJob := heap.Pop(&pq).(*Node)
			workers[freeWorkerIdx] = &Worker{
				timeLeft: getEstimatedTime(curJob),
				task:     curJob,
			}
			continue
		}

		updateWorkers(workers)

		// scan finished workers
		for idx := range workers {

			if workers[idx].timeLeft == 0 && workers[idx].task != nil {

				last_finished_iter = iter
				result = append(result, workers[idx].task.id)

				for j := range workers[idx].task.next {
					nextTask := workers[idx].task.next[j]
					if nextTask == nil {
						continue
					}
					nextTask.refCnt--
					if nextTask.refCnt == 0 {
						heap.Push(&pq, nextTask)
					}
				}

				// task is done
				workers[idx].task = nil

			}

		}

		iter++

	}

	return last_finished_iter + 1
}

func main() {
	content := getStdin()

	nodesMap := parseInput(content)
	nodesMapCloned := parseInput(content)

	fmt.Println("Serial sequence:", serialExecution(findOrigins(nodesMap)))
	fmt.Println("Parallel sequence:", parallelExecution(findOrigins(nodesMapCloned), 5))
}
