#$(Get-Date -UFormat "%A %m/%d/%Y %R %Z")
time := $(shell date +%s)
hash := $(shell git rev-parse HEAD)
test_file := "test.bf"
build: main.go
	@echo "building"
	go build -ldflags="-X main.BuildTime=$(time) -X main.CommitHash=$(hash)" -o vBrainFuck.exe

run: build
	./vBrainFuck

test: test.bf build
	./vBrainFuck -file $(test_file)
c-code: test.bf build
	./vBrainFuck -file $(test_file) > out/main.c
	gcc out/main.c -o ./out/a.exe
	./out/a.exe

clean:
	rm ./out/*
default: build