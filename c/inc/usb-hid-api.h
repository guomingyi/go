#ifndef __usb_hid_api_h__
#define __usb_hid_api_h__

// go function.
extern void go_print(int i);

int usb_init(void);
int usb_enumerate(int pid, int vid);
int usb_write(char *data);
int usb_read(char *data);
int usb_close(void);

#endif

