build-vmkit:
	go build -o ./bin/vmkit ./cmd/vmkit

build-avfvm:
	go build -o ./bin/avfvm ./cmd/avfvm

codesign-avfvm:
	codesign --entitlements ./res/avfvm.entitlements -s - ./bin/avfvm

build-and-codesign-all: build-vmkit build-avfvm codesign-avfvm
