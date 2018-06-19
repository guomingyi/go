#include <stdio.h>
#include "usb.h"

// golang call
int usb_init(void)
{
	printf("[%s] %s:%d\n", __FILE__, __func__, __LINE__);
	go_print(100);
    return usbdev_init();
}

// golang call
int usb_enumerate(int pid, int vid)
{
    return usbdev_enumerate(pid, vid);
}

// golang call
int usb_write(char *data)
{
   return usb_msg_out(data);
}

// golang call
int usb_read(char *data)
{
   return usb_msg_in(data);
}

// golang call
int usb_close(void)
{
    return usbdev_close();
}