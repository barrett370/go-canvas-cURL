.PHONY build:
build:
	go build -v -ldflags "-X main.authToken=${CANVAS_TOKEN}" main/Scrape.go

