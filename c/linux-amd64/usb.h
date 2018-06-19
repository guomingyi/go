#ifndef __usb_h__
#define __usb_h__

#include "../../c/inc/usb-hid-api.h"

int usbdev_init(void);
int usbdev_enumerate(int pid, int vid);
int usb_msg_out(char *buf);
int usb_msg_in(char *buf);
int usbdev_close(void);

#endif