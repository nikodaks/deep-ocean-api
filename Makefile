.PHONY:
.SILENT:
.DEFAULT_GOAL := run

install: 
	go-get
run:
	air
build:
	go build -o ./deep_ocean ./cmd/app/.
