package main

import (
	"fmt"
	"myProject/pkg/storage"
	"myProject/pkg/student"
)

func main() {
	MapStudent := make(map[string]*storage.Student, 0)
	fmt.Println("Введите имя, возраст и курс студента через пробел:")
	student.StudentPlay(MapStudent)
	fmt.Println("Вывожу всех введённых студентов:")
	for _, value := range MapStudent {
		fmt.Println(value.Info())
	}
}
