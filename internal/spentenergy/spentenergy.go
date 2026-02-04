package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("Ошибка входных начений")
	}
	m := MeanSpeed(steps, height, duration)
	d := duration.Minutes()
	calories := (weight * m * d) / minInH
	walkCalor := calories * walkingCaloriesCoefficient

	return walkCalor, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 || weight <= 0 || height <= 0 {
		return 0, errors.New("Ошибка входных начений")
	}
	m := MeanSpeed(steps, height, duration)
	d := duration.Minutes()
	calories := (weight * m * d) / minInH
	return calories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if steps <= 0 || duration <= 0 {
		return 0
	}

	d := Distance(steps, height)
	midSpeed := d / duration.Hours()
	return midSpeed

}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию

	stepsLong := height * stepLengthCoefficient
	distance := float64(steps) * stepsLong
	distInKm := distance / mInKm
	return distInKm
}
