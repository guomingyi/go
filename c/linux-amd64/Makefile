TARGET := libhello.a
BUILD := ../lib

SRCS := hello.c
SRCS += test.c

OBJS := $(SRCS:%.c=%.o)
CC := gcc
AR := ar
CFLAG := rcs

all:
	@rm -rf *.a *.o
	$(CC) -c $(SRCS)
	$(AR) $(CFLAG) $(TARGET) $(OBJS)
	@mv $(TARGET) $(BUILD)

clean:
	@rm -rf *.a *.o $(BUILD)/*.a

