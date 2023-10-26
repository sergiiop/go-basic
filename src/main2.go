package main

import (
	"fmt"
	"time"
)

type Interval struct {
	FechaInicio time.Time `json:"fecha_inicio"`
	FechaFin    time.Time `json:"fecha_fin"`
} 

func generarIntervalosMensuales(fechaInicial, fechaFinal time.Time) ([]Interval, error) {
	var intervalos []Interval

	for fechaInicial.Before(fechaFinal) || fechaInicial.Equal(fechaFinal) {
		fechaFinalDelMes := time.Date(fechaInicial.Year(), fechaInicial.Month()+1, 1, 0, 0, 0, 0, time.UTC).Add(-time.Second)
		if fechaFinalDelMes.After(fechaFinal) {
			fechaFinalDelMes = fechaFinal
		}

		intervalo := Interval{
			FechaInicio: fechaInicial,
			FechaFin:    fechaFinalDelMes,
		}

		intervalos = append(intervalos, intervalo)

		fechaInicial = fechaFinalDelMes.Add(time.Second)
	}

	return intervalos, nil
}

func main() {
	fechaInicial, _ := time.Parse("2006-01-02", "2023-01-15")
	fechaFinal, _ := time.Parse("2006-01-02", "2023-05-20")

	intervalos, err := generarIntervalosMensuales(fechaInicial, fechaFinal)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Imprime cada intervalo
	for _, intervalo := range intervalos {
		fmt.Printf("FechaInicio: %s, FechaFin: %s\n", intervalo.FechaInicio.Format("2006-01-02"), intervalo.FechaFin.Format("2006-01-02"))
	}
}
