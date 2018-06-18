package main

// #include "../c/inc/hello.h"
// #cgo LDFLAGS: ../c/lib/libhello.a -lm -ldl
// #cgo pkg-config: gtk+-3.0
import "C"

import (
	"fmt"
	"github.com/andlabs/ui"
)

//export go_print
func go_print(i C.int) {
	fmt.Printf("go print i = %v\n", uint32(i))
}

func main() {
	C.hello(9)
	C.test(6)
	fmt.Printf("Hello, world.\n")
	show_ui()
}

func show_ui() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter your name:"), false)
		box.Append(input, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 600, 400, false)
		window.SetMargined(true)
		window.SetChild(box)
		button.OnClicked(func(*ui.Button) {
			greeting.SetText("Hello, " + input.Text() + "!")
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
