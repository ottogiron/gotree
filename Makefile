MONGO_TEST_HOSTS ?= $(shell docker-machine ip default)
SOURCE_DIRS := $(shell go list ./... | grep -v /vendor/)

test: clean
	@docker-compose up -d
	MONGO_TEST_HOSTS=$(MONGO_TEST_HOSTS) \
	go test $(SOURCE_DIRS)


clean:
	@docker-compose stop
	@docker-compose rm -f
