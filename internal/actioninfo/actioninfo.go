package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for i := 0; i < len(dataset); i++ {
		err := dp.Parse(dataset[i])
		if err != nil {
			log.Printf("Ошибка выполнения %v", err)
			continue
		}
		s, errA := dp.ActionInfo()
		if errA != nil {
			log.Printf("Ошибка выполнения %v", err)
			continue
		}
		fmt.Print(s)

	}

}
