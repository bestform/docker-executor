.PHONY: build

build:
	GOOS=linux CGO_ENABLED=0 go build -o dockerExecutor main.go
	docker build -t docker-executor:latest .

run-example: build
	docker run -p 8080:8080 -e DOCKER_EXECUTOR_CMD="php -v" -e DOCKER_EXECUTOR_IDENTIFIER=php -v /var/run/docker.sock:/var/run/docker.sock docker-executor:latest
