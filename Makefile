.PHONY: build
build:
	go build -o ./bin/vmkit .

.PHONY: clean
clean:
	rm -rf ./bin
