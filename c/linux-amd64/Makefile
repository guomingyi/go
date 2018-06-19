# static:
#// #include "../c/inc/hello.h"
#// #cgo LDFLAGS: ../c/lib/libhello.a -lm -ldl

# shared:
#// #include "../c/inc/hello.h"
#// #cgo LDFLAGS: -L../c/lib -lhello
# export LD_LIBRARY_PATH=c/lib/

#gcc -shared -fPIC -o libhello.so usb.c -lusb-1.0 -I/usr/include/libusb-1.0 -std=c99 -L/usr/lib

SRCS :=
SRCS += usb-hid-api.c
SRCS += usb.c

TARGET := libusb-hid-api.so
BUILD := ../../bin

LDIR = /usr/lib
LFLAG = -lusb-1.0
INCLUDES = /usr/include/libusb-1.0 
CFLAGS = $(LFLAG) -I$(INCLUDES) -std=c99 -L$(LDIR)

OBJS := $(SRCS:%.c=%.o)
CC := gcc
AR := ar
CFLAG := rv 

all:
	@rm -rf *.a *.o
	$(CC) -shared -fPIC -o $(TARGET) $(SRCS) $(CFLAGS)
	@[ -d ${BUILD} ] || mkdir -p ${BUILD}
	@cp $(TARGET) $(BUILD)/

clean:
	@rm -rf *.a *.o *.so $(BUILD)/$(TARGET)

