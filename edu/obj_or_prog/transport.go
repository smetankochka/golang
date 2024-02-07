package main

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
}

type Car struct {
	Type     string
	Speed    float64
	FuelType string
}

type Motorcycle struct {
	Type     string
	Speed    float64
	FuelType string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	result := make(map[string]float64)
	for _, vehicle := range vehicles {
		val, ok := vehicle.(Car)
		if ok {
			result[val.Type] = vehicle.CalculateTravelTime(distance)
		} else {
			val, _ := vehicle.(Motorcycle)
			result[val.Type] = vehicle.CalculateTravelTime(distance)
		}
	}
	return result
}
