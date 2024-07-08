all: install build

install:
	@echo
	@echo "********** installing dependencies **********"
	go get ./...
	@echo
	@echo "********** all dependencies installed successfully **********"

build:
	@echo
	@echo "********** building executable **********"
	chmod +x build_local.sh
	./build_local.sh
	@echo
	@echo "********** clread has been built successfully **********"
