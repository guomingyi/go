package main

// #include "../c/inc/usb-hid-api.h"
// #cgo LDFLAGS: -L ../bin -lusb-hid-api
import "C"

import (
	"fmt"
	"github.com/andlabs/ui"
	"time"
	"unsafe"
)

//export go_print
func go_print(i C.int) {
	fmt.Printf("go print i = %v\n", uint32(i))
}

const PID int = 0x0001
const VID int = 0x534c

//var greeting *Label

func usb_dev_init() int {
	var ret int = (int)(C.usb_init())
	if ret == 0 {
		ret = (int)(C.usb_enumerate(C.int(PID), C.int(VID)))
		fmt.Printf("usb_enumerate ret:%d\n", ret)
	}
	return ret
}

func go_usb_write(msg string) {
	if msg != "" && len(msg) <= 64 {
		var buffer []byte
		buffer = []byte(msg)
		C.usb_write((*C.char)(unsafe.Pointer(&buffer[0])))
		fmt.Printf("go usb_write: %s\n", string(buffer[:]))
	}
}

func go_usb_read() string {
	var buffer [64]byte
	C.usb_read((*C.char)(unsafe.Pointer(&buffer[0])))
	return string(buffer[:])
}

func main() {
	// init usb hid device.
	if usb_dev_init() < 0 {
		fmt.Printf("usb_dev_init: failed!\n")
		return
	}
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
			go_usb_write(input.Text())
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

		go counter(greeting)
		time.Sleep(1000 * 1000)
	})
	if err != nil {
		panic(err)
	}
}

func counter(label *ui.Label) {
	for {
		rev := go_usb_read()
		ui.QueueMain(func() {
			label.SetText("receiver from usbhid: " + rev)
		})
		time.Sleep(1000 * 1000)
	}
}
