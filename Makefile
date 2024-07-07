all: install build

install:
	go get ./...
	echo "all dependencies installed successfully"

build:
	chmod +x build_local.sh
	./build_local.sh
	echo "clread has been built successfully"
