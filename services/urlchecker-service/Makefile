all: build run 

# Local build
build:
	/usr/local/go/bin/go get -d -v ./...
	/usr/local/go/bin/go build -o bin/urlchecker *.go
	chmod +x bin/urlchecker

run:
	go run *.go

compile:
	echo "Compiling for linux-amd64"
	GOOS=linux GOARCH=amd64 go build -o bin/urlchecker-linux-amd64 *.go

clean:
	rm -rf bin/urlchecker*

test:
	go test

# Docker dev build and test
docker-build-dev:
	docker build -t mrsouliner/urlchecker-dev:latest -t mrsouliner/urlchecker-dev:$(GIT_COMMIT_SHA)  -f Dockerfile.dev .

docker-run-dev:
	docker run urlchecker-dev:latest -d 

docker-run-test:
	docker run --rm --name urlchecker urlchecker-dev:latest go test

# Final build
docker-compile:
	docker run --env GOOS=linux --env GOARCH=amd64 --rm --name urlchecker urlchecker-dev:latest go build -o bin/urlchecker-linux-amd64

docker-build-final:
	docker build -t mrsouliner/urlchecker:latest -t mrsouliner/urlchecker:$(GIT_COMMIT_SHA) . 

docker-push-dev: 
	docker login -u="$(DOCKER_USER)" -p="$(DOCKER_PASS)"
	docker push mrsouliner/urlchecker-dev:latest
	docker push mrsouliner/urlchecker-dev:$(GIT_COMMIT_SHA)
