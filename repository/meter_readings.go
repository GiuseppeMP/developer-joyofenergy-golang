package repository

import (
	"joi-energy-golang/domain"
	"math/rand/v2"
	"time"
)

type MeterReadings struct {
	meterAssociatedReadings map[string][]domain.ElectricityReading
}

func NewMeterReadings(meterAssociatedReadings map[string][]domain.ElectricityReading) MeterReadings {
	return MeterReadings{meterAssociatedReadings: meterAssociatedReadings}
}

func (m *MeterReadings) GetReadings(smartMeterId string) []domain.ElectricityReading {
	v, ok := m.meterAssociatedReadings[smartMeterId]
	if !ok {
		return nil
	}
	return v
}

func (m *MeterReadings) StoreReadings(smartMeterId string, electricityReadings []domain.ElectricityReading) {
	m.meterAssociatedReadings[smartMeterId] = append(m.meterAssociatedReadings[smartMeterId], electricityReadings...)
}

func (m *MeterReadings) GetAllReadings() []domain.ElectricityReading {
	readings := make([]domain.ElectricityReading, 0)
	for i := 0; i < 150; i++ {
		readingTime := time.Now().AddDate(0, 0, -rand.IntN(365))
		reading := domain.ElectricityReading{Time: readingTime, Reading: rand.Float64()}
		readings = append(readings, reading)
	}
	return readings
}
func (m *MeterReadings) GetLastWeekReadings() float64 {
	readings := m.GetAllReadings()
	var sum float64
	for _, reading := range readings {
		if reading.Time.After(time.Now().AddDate(0, 0, -7)) {
			sum += reading.Reading
		}
	}
	return sum
}
