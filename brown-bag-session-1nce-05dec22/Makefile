.PHONY: all
all: flash test

.PHONY: fumpt-install
fumpt-install:
ifeq (, $(shell which gofumpt))
	go install mvdan.cc/gofumpt@v0.3.1
endif

.PHONY: fumpt
fumpt: fumpt-install
	gofumpt -l -w .

.PHONY: tidy
tidy:
	cd avrmodem && go mod verify &&	go mod tidy

.PHONY: build
build:
	cd avrmodem && tinygo build -target arduino-mega2560 -o bin/firmware firmware.go

.PHONY: flash
flash:
	cd avrmodem && tinygo flash -target arduino-mega2560 firmware.go

.PHONY: hello
hello:
	cd helloworld && tinygo build -target arduino-mega2560 hello.go

.PHONY: buzzer
buzzer:
	cd buzzer_demo && tinygo flash -target arduino-mega2560 make_noise.go

.PHONY: test
test:
	AT_TRACE=true avrmodem/bin/eeagent -c avrmodem/eeagent.conf.yaml

.PHONY: clean
clean:
	rm -f avrmodem/bin/firmware avrmodem/eeagent.db

.PHONY: at
at:
	# make dev=/dev/cu.usbmodem14401 cmd=+QGMR at
	cd sendat && go run sendat.go -v -d $(dev) -c $(cmd)
