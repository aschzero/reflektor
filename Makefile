IMAGE=aschzero/reflektor:latest
PWD=$(shell pwd)
CMD_DIR=$(PWD)/cmd/reflektor

default: build

release: image push

build:
	cd $(CMD_DIR) && go build -o $(PWD)/reflektor

image:
	docker build -t $(IMAGE) .

push:
	docker push $(IMAGE)