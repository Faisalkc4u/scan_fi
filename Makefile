APP_NAME=scan-fi
CONTAINER_NAME=scan-fi-app

.PHONY: build docker-build docker-run docker-stop docker-rm docker-restart

build:
	GOOS=linux GOARCH=amd64 go build -o main .

docker-build: build
	docker build -t $(APP_NAME) .

docker-run:
	docker run -d --name $(CONTAINER_NAME) --env-file .env -e ENV=production -p 8000:8000 $(APP_NAME)

docker-stop:
	-docker stop $(CONTAINER_NAME)

docker-rm:
	-docker rm $(CONTAINER_NAME)

docker-restart: docker-stop docker-rm docker-run
