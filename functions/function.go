package functions
import "fmt"

// No parameters, no return value
func PrintHello() {
    fmt.Println("Hello")
}
// Called like this:

// One parameter, one return value
func Hello(name string) string {
  return "Hello " + name
}

// Multiple parameters, multiple return values
func SumAndMultiply(a, b int) (int, int) {
    return a+b, a*b
}
// Called like this:
// aplusb, atimesb := SumAndMultiply(a, b)