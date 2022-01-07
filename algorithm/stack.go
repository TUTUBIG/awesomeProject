package main

type Stack []interface{}

func (s *Stack) Push(data ...interface{}) {
	*s = append(*s, data...)
}

func (s *Stack) Pop() (data interface{}) {
	l := len(*s)
	if l < 1 {
		return nil
	}
	d := (*s)[l-1]
	*s = (*s)[:l-1]
	return d
}
