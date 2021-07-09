package main

type Stack struct {
	slice []int
}

func (s *Stack) len() int {
	return len(s.slice)
}

func (s *Stack) push(val int) {
	s.slice = append(s.slice, val)
}

func (s *Stack) pop() int {
	if s.len() == 0 {
		panic("Error")
	}
	var tmp int = s.slice[s.len()-1]
	s.slice = s.slice[0 : s.len()-1]
	return tmp
}

func main() {

}
