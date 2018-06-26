package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"time"
)

func Show_ui() {
	fmt.Printf("show ui")

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
			Usb_write(input.Text())
			if input.Text() != "" {
				greeting.SetText(input.Text())
			} else {
				greeting.SetText("no cmd..")
			}
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		window.Show()
		
		go UpdateUI(greeting)
		time.Sleep(1000 * 1000)
	})
	if err != nil {
		panic(err)
	}
}

func UpdateUI(label *ui.Label) {
	for {
		ret, rev := Usb_read()
		if ret < 0 {
			fmt.Printf("Usb_read err: %d\n", ret)
		}
		ui.QueueMain(func() {
			label.SetText(rev)
		})
		time.Sleep(1000 * 1000)
	}
}








