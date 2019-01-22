all: build

build:
	go build -o bin/eve cmd/eve/*.go

robotgo:
	sudo apt-get install -y gcc libc6-dev
	sudo apt-get install -y libx11-dev xorg-dev libxtst-dev libpng++-dev
	sudo apt-get install -y xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
	sudo apt-get install -y libxkbcommon-dev
	sudo apt-get install -y xsel xclip

tools:
	go get -u github.com/golang/dep/cmd/dep

devtools:
	go get github.com/vektra/mockery/.../

deps:
	dep ensure --vendor-only

clean:
	rm -rf bin