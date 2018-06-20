# go build

SRCS := *.go
TARGET := factory_test
release := ../bin

FLAG += -v
FLAG += -work
# FLAG += -x

ARCH_FLAG := CGO_ENABLED=1 GOOS=linux GOARCH=amd64
#ARCH_FLAG := CGO_ENABLED=1 GOOS=windows GOARCH=amd64
#ARCH_FLAG := CGO_ENABLED=1 GOOS=darwin GOARCH=amd64


all:
	$(ARCH_FLAG) go build -o $(TARGET) $(FLAG) $(SRCS)
	@[ -d ${release} ] || mkdir -p ${release}
	@mv $(TARGET) $(release)
	@tree $(release)

clean:
	@rm -rf $(release)

