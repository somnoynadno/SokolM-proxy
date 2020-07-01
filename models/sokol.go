package models

import "time"

type SokolM struct {
	Time int64     `gorm:"PRIMARY_KEY" json:"time"`
	Date time.Time `gorm:"-"`
	EVS  float64   // EVS
	UVI  float64   // UVI
	L    float64   // L
	LI   float64   // LI
	RSSI float64   // Уровень сигнала GSM
	RN   float64   // Осадки, мм
	T    float64   // Температура, град. C
	WD   float64   // Направление ветра, град.
	HM   float64   // Влажность, %
	WV   float64   // Скорость ветра, м/с
	WM   float64   // Порыв ветра, м/с
	UV   float64   // Уровень ультрафиолета, Вт/м2
	Upow float64   // Напряжение питания, В
	PR1  float64   // Атм.давление, мм.рт.ст.
	PR   float64   // Атмосферное давление, гПа
	KS   float64   // KS
	TR   float64   // Счетчик
	TD   float64   // Точка росы, ⁰C
}