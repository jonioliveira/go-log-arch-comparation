.DEFAULT_GOAL := help

## General

# target: help - Display available recipes.
.PHONY: help
help:
	@egrep "^# target:" [Mm]akefile

# target: run singleton example
.PHONY: singleton
singleton:
	go run cmd/singleton/main.go


# target: run injection example
.PHONY: injection
injection:
	go run cmd/injection/main.go
