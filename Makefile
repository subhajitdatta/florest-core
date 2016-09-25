## simple makefile to log workflow
.PHONY: deploy

LOGLEVEL ?= 1
SWAGGER ?= 2
NEWAPP ?=3
GOFLAGS ?= $(GOFLAGS:)

PWD = $(shell pwd)
export GOPATH = $(shell echo $$GOPATH):$(PWD)/_libs:$(PWD)
export GOBIN = $(PWD)/bin
export GOROOT = $(shell echo $$GOROOT)


pkg: deploy
	@mv bin/`basename "$(CURDIR)"` bin/floRest_tmp
	@mv bin/floRest_tmp bin/floRest
	@tar zcvf deploy/floRest_service.tar.gz bin/

deploy: clean format build install conf swagger

build:
	@rm -rf pkg/ 2>/dev/null
	@rm -rf _libs/pkg/ 2>/dev/null
	@go build $(GOFLAGS) ./...

conf:
	@mkdir -p bin/conf
	@cp config/florest-core/conf.json bin/conf/conf.json
	@cp config/logger/logger.json bin/conf/logger.json
	@bash scripts/logger.sh $(LOGLEVEL) bin/conf/logger.json

install:
	@go get ./...

test: format clean install
	@ginkgo -r -v=true -cover=true ./src/test/
	@go test ./...
	
coverage: install
	@sh src/test/coverage.sh

bench: install
	@go test -run=NONE -bench=. src/test/perftest/*.go

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

newapp: copyNewAppSrc copyNewAppLibs renameNewAppDir

copyNewAppSrc:
ifneq ($(NEWAPP),3)
	@echo building new app
	@cp -r -f _newApp/* $(NEWAPP)
endif

copyNewAppLibs:
ifneq ($(NEWAPP),3)
	@echo copying dependent libs
	@cp -r -f _libs/* $(NEWAPP)/_libs/
	@mkdir -p $(NEWAPP)/_libs/src/github.com/jabong/florest-core
	@cp -r -f src $(NEWAPP)/_libs/src/github.com/jabong/florest-core/
endif

renameNewAppDir:
	@bash scripts/newapp.sh $(NEWAPP)

coverall:
	@go test -c -covermode=count -coverpkg ./...
	@mv floRest.test bin/floRest.test
	
## EOF
