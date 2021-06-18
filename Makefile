#$(Get-Date -UFormat "%A %m/%d/%Y %R %Z")

build: main.go
	@echo "building"
	go build -ldflags="-X main.BuildTime=now" -o vBrainFuck.exe

run: build
	./vBrainFuck

test: test.bf build
	./vBrainFuck -file test.bf
c-code: test.bf build
	./vBrainFuck -file test.bf > out/main.c
	gcc out/main.c -o ./out/a.exe
	./out/a.exe

clean:
	rm ./out/*
default: build