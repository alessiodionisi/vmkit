build-vm:
	go build -o ./bin/vm ./cmd/vm

codesign-vm:
	codesign --entitlements ./cmd/vm/vm.entitlements -s - ./bin/vm

build-and-codesign-vm: build-vm codesign-vm
