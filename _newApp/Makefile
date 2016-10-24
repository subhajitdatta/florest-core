## simple makefile to log workflow
.PHONY: deploy

LOGLEVEL ?= 1
SWAGGER ?= 2
GOFLAGS ?= $(GOFLAGS:)
PWD = $(shell pwd)
export GOPATH =$(PWD)/_libs:$(PWD)
export GOBIN = $(PWD)/bin
	
pkg: deploy
	@tar zcvf deploy/`basename "$(CURDIR)"`_service.tar.gz bin/

deploy: clean format build install conf swagger

build:
	@rm -rf pkg/ 2>/dev/null
	@rm -rf _libs/pkg/ 2>/dev/null
	@go build $(GOFLAGS) ./...

conf:
	@mkdir -p bin/conf
	@cp `find config/ -name "*.json"` bin/conf/
	@bash scripts/logger.sh $(LOGLEVEL) bin/conf/logger.json

install: 
	@go install ./...

test: format clean
	@ginkgo -r -v=true -cover=true ./src/

coverage:
	@sh src/test/coverage.sh

bench:
	@go test -run=NONE -bench=. $(GOFLAGS) ./...

clean:
	@go clean $(GOFLAGS) -i ./...
	
format:
	@go fmt $(GOFLAGS) ./...
	
codeanalysis:
	@go tool vet src/	
	
swagger:
ifneq ($(SWAGGER),2)
	@echo building swagger support
	@bash scripts/swagger.sh $(SWAGGER)
endif

## EOF
