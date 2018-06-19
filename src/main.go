package main

// #include "../c/inc/usb-hid-api.h"
// #cgo LDFLAGS: -L ../bin -lusb-hid-api
import "C"

import (
	"fmt"
	"github.com/andlabs/ui"
	"strconv"
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

func go_usb_write(msg string) int {
	var ret int = 0

	if msg != "" && len(msg) <= 64 {
		var buffer []byte
		buffer = []byte(msg)
		ret = (int)(C.usb_write((*C.char)(unsafe.Pointer(&buffer[0]))))
		if ret < 0 {
			fmt.Printf("go usb_write: err!!!\n")
		} else {
			fmt.Printf("go usb_write: %s success\n", string(buffer[:]))
		}
	}
	return ret
}

func go_usb_read() (int, string) {
	var buffer [64]byte
	ret := (int)(C.usb_read((*C.char)(unsafe.Pointer(&buffer[0]))))
	return ret, string(buffer[:])
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
	var init_retry int = 5

	for {
		ret, rev := go_usb_read()
		if ret < 0 {
			C.usb_close()
			time.Sleep(1000 * 1000 * 5000)
			init_retry--
			rev = "usb read err, retry:" + strconv.Itoa(init_retry) + " pls wait.."
			usb_dev_init()
			if init_retry == 0 {
				break
			}
		} else {
			if init_retry < 5 && init_retry > 0 {
				rev = "retry success:%d\n" + strconv.Itoa(init_retry)
				init_retry = 5
			}
		}

		ui.QueueMain(func() {
			label.SetText(rev)
		})
		time.Sleep(1000 * 1000)
	}

	if init_retry == 0 {
		ui.QueueMain(func() {
			label.SetText("retry failed, exit.")
		})
	}
}
