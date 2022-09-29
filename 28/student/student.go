package student

import (
	"Golang/storage"
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

func NewStudent() {
	m := make(map[string]*storage.Student)
	EOF := make(chan os.Signal, 1)
	signal.Notify(EOF, os.Interrupt) // через ctrl+c при использовании sigint завершает без подтверждения через enter
	defer func() {
		fmt.Println("Студенты из хранилища:")
		for _, v := range m {
			fmt.Println(v)
		}
	}()
	for {
		fmt.Println("Введите данные о студенте: Имя,Возраст,Курс:")
		scaner := bufio.NewScanner(os.Stdin)
		scaner.Scan()
		nameOfStudent := scaner.Text()
		s := strings.Fields(nameOfStudent)
		newName, age, grade := s[0], s[1], s[2]
		newAge, err := strconv.Atoi(age)
		if err != nil {
			panic(err)
		}
		newGrade, err := strconv.Atoi(grade)
		if err != nil {
			panic(err)
		}
		m[nameOfStudent] = &storage.Student{
			Name:  newName,
			Age:   newAge,
			Grade: newGrade,
		}

	}

}
