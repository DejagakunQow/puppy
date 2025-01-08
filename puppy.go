package puppy

import (
	"fmt"

	"github.com/DejagakunQow/dog"
)

func Bark() string {
	return "woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())

}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func from13() {
	fmt.Println("i'm from version 1.3.0")
}

func from14() {
	fmt.Println("i'm from version 1.4.0")
}
