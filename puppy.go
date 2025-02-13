package puppy

import (
	"github.com/palgunatm66/dog"
)

func Bark() string{
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func BigBark() string {
	return main.WhenGrownUp(Bark())
}

func BigBarks() string {
	return main.WhenGrownUp(Barks())
}

