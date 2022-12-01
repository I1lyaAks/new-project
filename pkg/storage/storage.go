package storage

import (
	"myProject/pkg/student"
	"strconv"
)

func (s student.Student) Info() string {
	return s.Name + " " + strconv.Itoa(s.Age) + " " + strconv.Itoa(s.Grade)
}
func (s *student.Student) Put(m map[string]*student.Student) {
	m[s.Name] = s
}
func (s student.Student) Get(m map[string]*student.Student) bool {
	_, x := m[s.Name]
	return x
}
