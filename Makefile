.PHONY: run
run:
	go build && ./goloree

.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build