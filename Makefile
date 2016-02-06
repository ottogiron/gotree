MONGO_TEST_HOSTS ?= $(shell docker-machine ip default)

test: clean
	@docker-compose up -d
	MONGO_TEST_HOSTS=$(MONGO_TEST_HOSTS) \
	 godep go test -v ./...


clean:
	@docker-compose stop
	@docker-compose rm -f
