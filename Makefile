include .env

.PHONY app:
app:
	cd main && go build -v -ldflags "-X main.authToken=${CANVAS_TOKEN}" -o ../bin/scrape

