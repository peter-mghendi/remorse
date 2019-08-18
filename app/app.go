package app

import (
	"fmt"
	"strings"

	"github.com/alwindoss/morse"
	"github.com/eiannone/keyboard"
)

// Application represents the application.
type Application struct {
}

// writeMorse outputs the provided rune in morse code.
func writeMorse(c rune) error {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader(string(c)))	
	if err != nil {
		return err
	}

	fmt.Printf("%s ", string(morseCode))
	return nil
}

// Init sets up the application.
func (a *Application) Init() {
	e := keyboard.Open()
	if e != nil {
		panic(e)
	}
}

// Run starts the Application.
func (a *Application) Run() {
	fmt.Println("Remorse is running.\nPress ESC to exit.\n")
	for {
		char, key, err := keyboard.GetKey()
		if (err != nil) {
			panic(err)
		} else if (key == keyboard.KeyEsc) {
			break
		} else if (key == keyboard.KeySpace) {
			fmt.Print("/ ")
			continue
		}else if (key == keyboard.KeyEnter) {
			fmt.Println()
			continue
		}
		writeMorse(char)
	}
}
