.PHONY: build release

build:
	go build -o forward *.go

release:
	GOOS=darwin GOARCH=amd64 go build -o forward-darwin-x64
	GOOS=linux GOARCH=amd64 go build -o forward-linux-x64
	GOOS=windows GOARCH=amd64 go build -o forward-windows-x64
	tar czvf forward-darwin-x64.tar.gz forward-darwin-x64 README.md LICENSE
	tar czvf forward-linux-x64.tar.gz forward-linux-x64 README.md LICENSE
	tar czvf forward-windows-x64.tar.gz forward-windows-x64 README.md LICENSE