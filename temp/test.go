package temp

import (
	"fmt"
)

func Hello() {
	fmt.Println("hello")
}

func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	} else {
		return "One for you, one for me."
	}
}

// func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
// 	return (float64(1/100) * float64(successRate)) * float64(productionRate)
// }

func CleanupMessage(oldMsg string) string {
	i := 0
	end := len(oldMsg) - 1
	for string(oldMsg[i]) == "*" || string(oldMsg[i]) == " " || string(oldMsg[i]) == "\n" {
		i++
	}
	fmt.Println(string(oldMsg[i]), ":", i)

	for string(oldMsg[end]) == "*" || string(oldMsg[end]) == " " || string(oldMsg[end]) == "\n" {
		end--
	}
	var msg string
	for i <= end {
		msg += string(oldMsg[i])
		i++
	}
	return msg

}
