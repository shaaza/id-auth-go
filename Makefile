.PHONY: all
all: build-deps build fmt vet lint test

GLIDE_NOVENDOR=$(shell glide novendor)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
APP_EXECUTABLE="out/micro-auth"

setup-glide:
	curl https://glide.sh/get | sh
setup:
	go get -u github.com/golang/lint/golint

build-deps:
	glide install

compile:
	mkdir -p out/
	go build -o $(APP_EXECUTABLE)

fmt:
	go fmt $(GLIDE_NOVENDOR)

vet:
	go vet $(GLIDE_NOVENDOR)

lint:
	@for p in $(GLIDE_NOVENDOR); do \
		echo "==> Linting $$p"; \
		golint $$p | grep -vwE "exported (function|method|type) \S+ should have comment" | true; \
	done

test:
	ENVIRONMENT=test go test $(GLIDE_NOVENDOR) -p=1 -v

coverage:
	ENVIRONMENT=test go test $(GLIDE_NOVENDOR) -p=1 -cover | grep coverage:

coverage-report:
	ENVIRONMENT=test gocov test $(GLIDE_NOVENDOR) | gocov-html > docs/coverage.html

cyclomatic-complexity:
	gocyclo $(shell glide novendor | xargs -n1 dirname | xargs -n1 basename | sed \$$d | xargs echo)

build: build-deps compile fmt vet lint

copy-config:
	cp application.yml.sample application.yml
