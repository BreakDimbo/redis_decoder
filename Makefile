export GOPATH := $(CURDIR)

decoder:
	@echo "Building redis decoder..."
	go build -o bin/rdcoder src/main.go