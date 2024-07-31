package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"golang.design/x/clipboard"
)

func main() {
	// Initialize clipboard
	if err := clipboard.Init(); err != nil {
		panic(err)
	}

	fmt.Println("Press 'Ctrl+Alt+Q' to paste clipboard data...")

	lastEventTime := time.Now()

	// Register a key event listener
	hook.Register(hook.KeyDown, []string{"ctrl", "alt", "q"}, func(e hook.Event) {
		now := time.Now()
		if now.Sub(lastEventTime) < 500*time.Millisecond {
			// Debounce: Ignore if the key was pressed recently
			return
		}
		lastEventTime = now

		fmt.Println("Keybinding detected! Reading clipboard...")

		// Read the clipboard
		time.Sleep(time.Second / 2) // wait some time, this will help to get all the data
		clipData := string(clipboard.Read(clipboard.FmtText))

		if clipData != "" {
			fmt.Println("Clipboard data:", clipData)

			// Ensure that typing is done properly
			for i, char := range clipData {
				fmt.Printf("Typing char '%c' (index %d)\n", char, i)
				err := robotgo.KeyTap(string(char))
				if err != nil {
					log.Printf("Error in typing %c, reason: %v \n", char, err.Error())
				}
				time.Sleep(50 * time.Millisecond) // Adjust typing speed as needed
			}
		} else {
			fmt.Println("Clipboard is empty.")
		}
	})
	// Start the event listener
	fmt.Println("Event listener started...")
	s := hook.Start()
	<-hook.Process(s)
}
