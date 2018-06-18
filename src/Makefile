NAME = main
BUILD = ../build

FLAG += -v
FLAG += -x
FLAG += -work

all:
	go build -o $(NAME) $(FLAG) *.go
	@[ -d ${BUILD} ] || mkdir -p ${BUILD}
	@mv $(NAME) ../build/
	@tree ../build/

clean:
	@rm -rf ../build

