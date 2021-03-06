BINARY := cronparser
TEST_CRON ?= "*/15 0 1,15 * 5-2 /usr/bin/find"

.PHONY: build
build:
	go build -o $(BINARY) *.go

.PHONY: test
test: build
	./$(BINARY) $(TEST_CRON)
