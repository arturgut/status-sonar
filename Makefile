### LOCAL BUILD ### 

all: build run 

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

### DOCKER ### 

# Docker dev build and test
docker-build-dev:
	export GIT_COMMIT_SHA=$(git log -n 1 --pretty=format:'%h')
	cd services/urlchecker-service/; docker build -t mrsouliner/urlchecker-dev:latest -t mrsouliner/urlchecker-dev:${GIT_COMMIT_SHA} -f Dockerfile.dev .

# Final build
docker-compile:
	cd services/urlchecker-service/ && docker run --env GOOS=linux --env GOARCH=amd64 --rm -v ${PWD}/bin:/go/src/app/bin mrsouliner/urlchecker-dev:latest go build -o ./bin/urlchecker-linux-amd64
docker-build:
	export GIT_COMMIT_SHA=$(git log -n 1 --pretty=format:'%h')
	cd services/urlchecker-service/; docker build -t mrsouliner/urlchecker:latest -t mrsouliner/urlchecker:${GIT_COMMIT_SHA} -f Dockerfile .

docker-push-dev:
	export GIT_COMMIT_SHA=$(git log -n 1 --pretty=format:'%h') 
	docker login -u="${DOCKER_USER}" -p="${DOCKER_PASS}"
	docker push mrsouliner/urlchecker-dev:latest
	docker push mrsouliner/urlchecker-dev:${GIT_COMMIT_SHA}

docker-push:
	export GIT_COMMIT_SHA=$(git log -n 1 --pretty=format:'%h')
	docker login -u="${DOCKER_USER}" -p="${DOCKER_PASS}"
	docker push mrsouliner/urlchecker:latest
	docker push mrsouliner/urlchecker:${GIT_COMMIT_SHA}
