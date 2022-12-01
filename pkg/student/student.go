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

func StudentPlay(m map[string]*storage.Student) {
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
		student := storage.Student{studentName, studentAge, studentGrade}
		if err := student.Get(m); err == false {
			student.Put(m)
		} else {
			fmt.Print("Студент с таким именем уже есть в хранилище! Попробуйте снова...\n")
		}
	}
}
