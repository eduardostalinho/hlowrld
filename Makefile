default: build-image

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hlowrld .

build-image: build
	docker build . -t hlowrld
