package interfaces

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitor_name string) string
}

type German struct{}

// Greet implements Greeter.
func (g German) Greet(visitor_name string) string {
	return fmt.Sprintf("Hallo %s!", visitor_name)
}

// LanguageName implements Greeter.
func (g German) LanguageName() string {
	return "German"
}



type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}


func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}


// // Main function to test
// func main() {
// 	germanGreeter := German{}
// 	italianGreeter := Italian{}

// 	fmt.Println(SayHello("Dietrich", germanGreeter)) // Output: I can speak German: Hallo Dietrich!
// 	fmt.Println(SayHello("Luca", italianGreeter))    // Output: I can speak Italian: Ciao Luca!
// }