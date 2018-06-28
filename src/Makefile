OUT := js-ftm

all: clean Linux Windows info

l: clean Linux info

w: clean Windows info

Linux:
	@[ -d build ] || mkdir build
	@./mk ${PWD} ${OUT} Linux amd64
	@cp script/download.sh build

Windows:
	@[ -d build ] || mkdir build
	@./mk ${PWD} ${OUT} Windows amd64 
	@cp script/download.bat build
	
info:
	@tree  build

clean:
	@./mk ${PWD} ${OUT} xx xx clean 


