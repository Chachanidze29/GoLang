package main

type Queue struct {
	slice []int
}

func (q *Queue) len() int {
	return len(q.slice)
}

func (q *Queue) Enqueue(val int) {
	q.slice = append(q.slice, val)
}

func (q *Queue) Dequeue() int {
	if q.len() == 0 {
		panic("Error")
	}
	var tmp int = q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return tmp
}

func main() {

}
