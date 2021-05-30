build-vmkit:
	go build -o ./bin/vmkit ./cmd/vmkit

build-vmkit-vm:
	go build -o ./bin/vmkit-vm ./cmd/vm

codesign-vmkit-vm:
	codesign --entitlements ./res/vmkit.entitlements -s - ./bin/vmkit-vm

build-and-codesign-all: build-vmkit build-vmkit-vm codesign-vmkit-vm
