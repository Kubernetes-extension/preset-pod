GOCMD=GOOS=linux go
PROJECT=extension
SERVICE=preset-pod
REGISTRY=xxx.xxx.xxx.xxx
TAG=latest
REVISION=v1.1
RELEASE=production
BUILD_HASH=${shell git rev-parse HEAD}

build_pro:
	go build -o bin/$(SERVICE) main.go

push_pro:
	docker build -f Dockerfile -t $(REGISTRY)/$(PROJECT)/$(SERVICE):$(REVISION) .
	docker push $(REGISTRY)/$(PROJECT)/$(SERVICE):$(REVISION)