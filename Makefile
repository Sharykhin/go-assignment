.PHONY: run test

run:
	go run cmd/server/main.go

test:
	go test -v -count=1 ./switcher