ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
OS := $(shell uname -s)
export PATH := bin:$(PATH)

start:
	docker-compose up -d

status:
	docker-compose ps

stop:
	docker-compose stop

restart:
	docker-compose restart && make status

install-mockgen:
	go install github.com/golang/mock/mockgen@73266f9

generate-mocks:
	mockgen -source=handler/handler.go -destination=handler/mocks/mock_interface.go -package=mocks

run-appointments_scheduler:
	set -o allexport && \
	source .env set && \
	go run appointments_scheduler.main.go
