all: build

.PHONY: build
build:
	@echo "Building CMS application"
	@go build -v .

.PHONY: run
run:
	@echo "Running CMS application"
	@./go-cms

.PHONY: build-run
build-run: build run
	
.PHONY: clean
clean: 
	@go clean -i