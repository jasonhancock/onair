all: build
	@true

build:
	GOOS=linux GOARCH=arm go build -o onair

push:
	rsync -av onair bikelights.jasonhancock.com:
	rsync -av onair.py bikelights.jasonhancock.com:
