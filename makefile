# Make is verbose in Linux. Make it silent.
# Also we can silent a command with putting @ at the bigining of a command
MAKEFLAGS += --silent

build:
	echo " >> Building binary..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-s' -o ./bin/screenshot
	echo " >> Done!"

build-container:
	echo " >> Building container..."
	docker build --tag screenshot:1.0 .
	echo " >> Done!"

run: 
	echo " >> Running..."
	docker container run --rm -p 8080:8080 $(timeout) screenshot:1.0

all:build build-container run

scale:build build-container
	docker-compose up --scale screenshot=$(instance)
