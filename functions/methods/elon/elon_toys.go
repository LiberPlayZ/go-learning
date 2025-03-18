package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery-car.batteryDrain >= 0 {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

func (car *Car) DisplayDistance() string {
	s := fmt.Sprintf("Driven %d meters", car.distance)
	return s
}

func (car *Car) DisplayBattery() string {
	s := fmt.Sprintf("Battery at %d%%", car.battery)
	return s
}

func (car *Car) CanFinish(trackDistance int) bool {
	for car.battery >= car.batteryDrain {
		if car.distance >= trackDistance {
			return true
		}
		car.Drive()
	}
	return car.distance >= trackDistance
}
