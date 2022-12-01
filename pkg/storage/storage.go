package storage

import (
	"strconv"
)

type Student struct {
	Name string

	Age int

	Grade int
}

func (s Student) Info() string {
	return s.Name + " " + strconv.Itoa(s.Age) + " " + strconv.Itoa(s.Grade)
}
func (s *Student) Put(m map[string]*Student) {
	m[s.Name] = s
}
func (s Student) Get(m map[string]*Student) bool {
	_, x := m[s.Name]
	return x
}
