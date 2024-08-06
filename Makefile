all: install build

install:
	@echo
	@echo "********** installing dependencies **********"
	go get .
	@echo
	@echo "********** all dependencies installed successfully **********"

build:
	@echo
	@echo "********** building executable **********"
	go build -o clread
	rm -f $(HOME)/go/bin/clread
	mv clread $(HOME)/go/bin/
	@echo
	@echo "********** clread has been built successfully **********"
