package main

import (
	"fmt"
	"github.com/andlabs/ui"
)

func main() {
	fmt.Printf(" test main!\n")
	show_ui()
}

func show_ui() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("Send cmd to usbhid")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Terminal:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello1", 600, 400, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {

		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
