package main

import (
	"fmt"
	"usbhid"
)

func main() {
	fmt.Printf(" test main!\n")
	Connect("hid")
}

func Connect(path string) int {
	devs := usbhid.HidEnumerate(0, 0)

	for _, dev := range devs {
		if match(&dev) {
			fmt.Printf("hidapi - connect - low level open\n")
			_, err := dev.Open()
			if err != nil {
				fmt.Printf("hidapi Open failed:\n", dev)
				break
			}
			fmt.Printf("hidapi - connect -done\n")
			return 0
		}
	}
	fmt.Printf("hidapi Connect failed.\n")
	return -1
}

func match(d *usbhid.HidDeviceInfo) bool {
	vid := d.VendorID
	pid := d.ProductID
	ret := vid == 0x534c && (pid == 0x0001)
	return ret
}
