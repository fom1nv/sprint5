package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	str := strings.Split(datastring, ",")
	if len(str) != 3 {
		return errors.New("Ошибка парсинга")
	}
	steps, errS := strconv.Atoi(str[0])
	if errS != nil {
		return errors.New("Нулевой индекс парсинга не соотвествует типу int")
	}
	if steps <= 0 {
		return errors.New("Неверно переданное значение шагов")
	}
	t.Steps = steps
	t.TrainingType = str[1]
	ti, errT := time.ParseDuration(str[2])
	if errT != nil {
		return errors.New("ошибка преобразования времени в строку")
	}
	if ti <= 0 {
		return errors.New("Неверно переданное значение времени")
	}
	t.Duration = ti
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	d := spentenergy.Distance(t.Steps, float64(t.Height))

	srSpeed := spentenergy.MeanSpeed(t.Steps, float64(t.Height), t.Duration)
	tip := t.TrainingType
	var cal float64
	var err error
	if tip == "Ходьба" {
		cal, err = spentenergy.WalkingSpentCalories(t.Steps, float64(t.Weight), float64(t.Height), t.Duration)
		if err != nil {
			return "", errors.New("Ошибка при вызове WalkCalories")
		}
	} else if tip == "Бег" {
		cal, err = spentenergy.RunningSpentCalories(t.Steps, float64(t.Weight), float64(t.Height), t.Duration)
		if err != nil {
			return "", errors.New("ошибка RunCalor")
		}
	} else {
		return "", errors.New("неизвестный тип тренировки")
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), d, srSpeed, cal), nil

}
