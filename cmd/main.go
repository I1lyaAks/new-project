package main

import (
	"fmt"
	"myProject/pkg/student"
)

func main() {
	MapStudent := make(map[string]*student.Student, 0)
	fmt.Println("Введите имя, возраст и курс студента через пробел:")
	StudentPlay()
	fmt.Println("Вывожу всех введённых студентов:")
	for _, value := range MapStudent {
		fmt.Println(value.Info())
	}
}
