package speed

import (
	"fmt"
	"testing"
)

func TestCanFinish(t *testing.T) {
	car := NewCar(5, 2)
	track := NewTrack(100)
	ans := CanFinish(car, track)
	fmt.Println("enter")
	if ans != true {
		t.Errorf("fail")
	}

}
