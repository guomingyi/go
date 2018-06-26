OUT := test

all: clean Linux Windows info

l: clean Linux info

w: clean Windows info

Linux:
	@[ -d build ] || mkdir build
	@./mk ${PWD} ${OUT} Linux amd64

Windows:
	@[ -d build ] || mkdir build
	@./mk ${PWD} ${OUT} Windows amd64 
	
info:
	@tree  build

clean:
	@./mk ${PWD} ${OUT} xx xx clean 


