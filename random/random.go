package chance
import ("math/rand"
        "time"
       )
// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	n := 1 + rand.Intn(20) // a ≤ n ≤ b
    return n
}
// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
    rand.Seed(time.Now().UnixNano())
    f:=rand.Float64()*12
    if f>12{
    	f=12
        }    
    return f 
}
// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant","fox","giraffe","hedgehog"}
    rand.Shuffle(len(animals), func(i, j int) {
		animals[i], animals[j] = animals[j], animals[i]
	})
    return animals
}