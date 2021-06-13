build-vmkit:
	go build -o ./bin/vmkit ./cmd/vmkit

build-vmkitd:
	go build -o ./bin/vmkitd ./cmd/vmkitd

build-avfvm:
	go build -o ./bin/avfvm ./cmd/avfvm

codesign-avfvm:
	codesign --entitlements ./res/avfvm.entitlements -s - ./bin/avfvm

build-and-codesign-all: build-vmkit build-vmkitd build-avfvm codesign-avfvm
