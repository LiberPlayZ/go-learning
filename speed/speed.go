package speed

// import (
// 	"fmt"
// )

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		batteryDrain: batteryDrain,
		speed:        speed,
		distance:     0,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		return Car{
			battery:      car.battery - car.batteryDrain,
			batteryDrain: car.batteryDrain,
			speed:        car.speed,
			distance:     car.distance + car.speed,
		}
	} else {
		return car
	}

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for car.battery > car.batteryDrain{
		car = Drive(car)
		if  car.distance >= track.distance{
			break
		}
	}
	if car.distance >= track.distance {
		return true
	} else {
		return false
	}
}
