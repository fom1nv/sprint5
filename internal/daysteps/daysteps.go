package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	str := strings.Split(datastring, ",")
	if len(str) != 2 {
		return errors.New("Ошибка парсинга Str Не равно 2")
	}
	steps, errS := strconv.Atoi(str[0])
	if errS != nil {
		return errors.New("Неверное значение шагов после парсинга")
	}
	if steps <= 0 {
		return errors.New("Неверное значение шагов после парсинга")
	}
	ds.Steps = steps
	t, errT := time.ParseDuration(str[1])
	if errT != nil {
		return errors.New("неверное значение времени после парсинга")
	}
	if t <= 0 {
		return errors.New("Неверное значение времени после парсинга")
	}
	ds.Duration = t
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию

	dist := spentenergy.Distance(ds.Steps, float64(ds.Height))

	cal, err := spentenergy.WalkingSpentCalories(ds.Steps, float64(ds.Weight), float64(ds.Height), ds.Duration)
	if err != nil {
		return "", errors.New("Ошибка вызова WakkingCalories в ActionInfo")
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, dist, cal), nil
}
