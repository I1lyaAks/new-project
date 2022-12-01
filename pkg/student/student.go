package student

import (
	"bufio"
	"fmt"
	"io"
	"myProject/pkg/storage"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	Name string

	Age int

	Grade int
}

func StudentPlay() {
	MapStudent := make(map[string]*Student, 0)
	var in = bufio.NewReader(os.Stdin)
	for i, err := in.ReadString('\n'); err != io.EOF; i, err = in.ReadString('\n') {
		sliseStruct := strings.Fields(i)

		if len(sliseStruct) < 3 {
			fmt.Println("что-то пошло не так... Пожалуйста, попробуйте ещё раз.")
			continue
		}
		studentName := sliseStruct[0]
		studentAge, errAge := strconv.Atoi(sliseStruct[1])
		studentGrade, errGrade := strconv.Atoi(sliseStruct[2])

		if errAge != nil || errGrade != nil {
			fmt.Print("Ошибка при обработке возраста студента или его номера курса ! Пожалуйста, попробуйте снова...\n")
			continue
		}
		student := Student{studentName, studentAge, studentGrade}
		if err := storage.Get(student, MapStudent); err == false {
			storage.Put(MapStudent)
		} else {
			fmt.Print("Студент с таким именем уже есть в хранилище! Попробуйте снова...\n")
		}
	}
}
