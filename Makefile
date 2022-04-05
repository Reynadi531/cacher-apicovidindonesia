APP_NAME=cacher-apicovidindonesia

build:
	go build -o bin/main main.go

run:
	go run .

docker.build:
	docker build . -t ghcr.io/reynadi531/$(APP_NAME):latest -f Dockerfile

docker.push:
	docker push ghcr.io/reynadi531/$(APP_NAME):latest

docker.run:
	docker run -v $(PWD)/results:/app/results ghcr.io/reynadi531/$(APP_NAME):latest
