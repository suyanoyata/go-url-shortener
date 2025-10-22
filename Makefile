APP_NAME := go-url-shortener

PLATFORM := $(shell uname -s | tr ' ' '-' | tr A-Z a-z)
schema := debug

.PHONY: all build clean run

all: build run

build:
	go build -o build/$(PLATFORM)/$(APP_NAME)

run:
	build/$(PLATFORM)/$(APP_NAME)

clean:
	rm -f build/$(PLATFORM)/$(APP_NAME)