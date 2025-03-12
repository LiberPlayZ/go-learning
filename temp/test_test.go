package temp

import (
	"fmt"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
// func TestHelloName(t *testing.T) {
// 	name := "Gladys"
// 	msg := ShareWith(name)
// 	fmt.Println(msg)
// 	if msg != "One for "+name+", one for me." {
// 		t.Fail()
// 	}

// }

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.


func TestCleanupMessage(t *testing.T) {
	msg :="***************************    BUY NOW, SAVE 10%   ***************************"
	ans := CleanupMessage(msg)
	fmt.Println("ans",msg)
	if ans != "BUY NOW, SAVE 10%" {
		t.Fail()
	}

}
