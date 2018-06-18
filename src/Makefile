# go build

SRCS := *.go
TARGET := main
release := ../build

FLAG += -v
FLAG += -work
# FLAG += -x

all:
	go build -o $(TARGET) $(FLAG) $(SRCS)
	@[ -d ${release} ] || mkdir -p ${release}
	@mv $(TARGET) $(release)
	@tree $(release)

clean:
	@rm -rf $(release)

