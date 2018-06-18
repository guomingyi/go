# c language code.
CPATH := ./c

# go language code.
GO := ./src

all:
	make -C $(CPATH)
	make -C $(GO)

clean:
	make clean -C $(CPATH)
	make clean -C $(GO)

