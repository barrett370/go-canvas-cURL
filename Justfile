linux:
	GOOS=linux go build -o ./bin/scrape-linux64


windows:
  GOOS=windows go build -o ./bin/scrape-win64.exe

all: linux windows
