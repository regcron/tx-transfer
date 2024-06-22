GO_CMD=go
GO_TEST=CGO_ENABLED=0 && $(GO_CMD) test
SHELL:=/usr/bin/env sh

.PHONY:unit-tests
unit-tests:
	go clean -testcache
	@-$(GO_TEST) `go list ./bounded_contexts/... | grep -v db | grep -v testing | grep -v injectors` -v -cover ||:

.PHONY:build
build:
	docker-compose build
.PHONY:run
run:
	docker-compose up -d

.PHONY:stop
stop:
	docker-compose stop

.PHONY:logs
logs:
	docker-compose logs -f tx