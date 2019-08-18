package app

import (
	"fmt"
	"strings"
	"time"

	"github.com/alwindoss/morse"
	"github.com/eiannone/keyboard"
	"github.com/inancgumus/screen"
)

// Application represents the application.
type Application struct {
	cache []string
}

// writeMorse outputs the provided rune in morse code.
func getMorse(c rune) (string, error) {
	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader(string(c)))
	if err != nil {
		return "", err
	}

	return string(morseCode), nil
}

// Init sets up the application.
func (a *Application) Init() {
	e := keyboard.Open()
	if e != nil {
		panic(e)
	}
}

// Run starts the Application.
func (a *Application) Run() error {
	fmt.Println("Remorse is starting up.\nPress ESC to exit.")
	time.Sleep(3 * time.Second)
	screen.Clear()
	screen.MoveTopLeft()

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			return err
		} else if key == keyboard.KeyEsc {
			break
		} else if key == keyboard.KeySpace {
			fmt.Print("/ ")
			continue
		} else if key == keyboard.KeyBackspace {
			if len(a.cache) <= 0 {
				continue
			}

			a.cache = a.cache[:len(a.cache)-1]
			screen.Clear()
			screen.MoveTopLeft()
			fmt.Printf("%s ", strings.Join(a.cache, " "))
			continue
		} else if key == keyboard.KeyEnter {
			fmt.Println()
			continue
		}

		str, err := getMorse(char)
		if err != nil {
			return err
		}

		a.cache = append(a.cache, str)
		fmt.Printf("%s ", str)
	}
	return nil
}
