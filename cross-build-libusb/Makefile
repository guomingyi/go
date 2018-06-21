OUT := usb-hidapi

all:
	@[ -d build ] || mkdir build
	@./mk ${PWD} ${OUT} amd64 all
	@cp -a build/ ~/share/

linux:
	@[ -d build ] || mkdir build
	@./mk ${PWD} ${OUT} amd64 l
	@cp -a build/ ~/share/

win:
	@[ -d build ] || mkdir build
	@./mk ${PWD} ${OUT} amd64 w
	@cp -a build/ ~/share/

clean:
	@rm -rf build


