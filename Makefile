all: build
	@true

build:
	GOOS=linux GOARCH=arm go build -o onair

push:
	rsync -av onair pi@onair.jasonhancock.com:
	rsync -av onair.py pi@onair.jasonhancock.com:
