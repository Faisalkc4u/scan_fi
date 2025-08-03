APP_NAME=scan-fi
CONTAINER_NAME=scan-fi-app
IMAGE_TAG=latest
FULL_IMAGE=$(APP_NAME):$(IMAGE_TAG)

.PHONY: build docker-build-prod docker-run docker-stop docker-rm docker-restart

build:
	GOOS=linux GOARCH=amd64 go build -o main .

docker-build-prod: build
	docker buildx build --platform linux/amd64 -t $(FULL_IMAGE) --load .

docker-run:
	docker run -d --name $(CONTAINER_NAME) --env-file .env -e ENV=production -p 8000:8000 $(FULL_IMAGE)

docker-stop:
	-docker stop $(CONTAINER_NAME)

docker-rm:
	-docker rm $(CONTAINER_NAME)

docker-restart: docker-stop docker-rm docker-run
