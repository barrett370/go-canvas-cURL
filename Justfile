
all: linux linuxARM windows mac macARM

linux:
	GOOS=linux go build -o ./bin/scrape-linux64

linuxARM: 
	GOOS=linux GOARCH=arm go build -o ./bin/scrape-linuxARM

windows:
  GOOS=windows go build -o ./bin/scrape-win64.exe

mac:
  GOOS=darwin go build -o ./bin/scrape-macos64
macARM:
  GOOS=darwin GOARCH=arm64 go build -o ./bin/scrape-macosarm
