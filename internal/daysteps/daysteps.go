package daysteps

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	stepLength = 0.65 // Длина одного шага в метрах
	mInKm      = 1000 // Количество метров в одном километре
)

// parsePackage парсит строку с данными о шагах и продолжительности прогулки
func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("должно быть две части")
	}
	stepsStr := parts[0]
	durationStr := parts[1]
	// Парсим количество шагов

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		log.Println(err.Error())
		return 0, 0, errors.New("неправильное количество шагов")
	}

	// Проверяем что шаги положительные
	if steps <= 0 {
		return 0, 0, errors.New("количество шагов должно быть больше нуля")
	}

	// Парсим продолжительность
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return 0, 0, err
	}

	// Проверяем что продолжительность положительная
	if duration <= 0 {
		return 0, 0, errors.New("продолжительность должна быть больше нуля")
	}

	return steps, duration, nil
}

// DayActionInfo возвращает информацию о прогулке
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {

		return ""
	}

	// Рассчитываем пройденное расстояние в км
	distanceMeters := float64(steps) * stepLength
	distanceKm := distanceMeters / mInKm

	// Рассчитываем сожженные калории
	calories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	result := fmt.Sprintf("Количество шагов: %d.\n"+
		"Дистанция составила %.2f км.\n"+
		"Вы сожгли %.2f ккал.\n",
		steps, distanceKm, calories)
	return result
}
