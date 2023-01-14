GOCMD			=	go
GOFMT			=	$(GOCMD) fmt
GOFUMPT		=	gofumpt
GOVET			=	$(GOCMD) vet
GOBUILD		=	$(GOCMD) build
GORETURNS	= goreturns
LINTER    = golangci-lint
APP       = todo-backend
PLATFORM  = linux/arm64
ENV       = local
PORT			= 8080
CONTAINER	= ko.local/$(APP)
PGURL     = postgres://user:pass@127.0.0,1:5432/localdb?sslmode=disable

all:
	make goreturns
	make vet
	make lint-fix
	make lint
	make test
	make swaggo
	make build-go

validate:
	$(GOFMT) -w -s ./...

goreturns:
	$(GORETURNS) -w .

lint:
	$(LINTER) run

lint-fix:
	$(LINTER) run --fix

vet:
	$(GOVET) ./...

swaggo:
	swag init

build:
	ko publish --local --platform=$(PLATFORM) --base-import-paths .

build-go:
	$(GOBUILD) main.go

test:
	$(GOCMD) test -v -cover -covermode=atomic ./...

run:
	docker run -td -i --network docker.internal --env-file config/$(ENV).env -p $(PORT):$(PORT) --name $(APP) $(CONTAINER)

run-go:
	$(GOCMD) run main.go

run-air:
	air

tbls:
	tbls doc "$(PGURL)"

clean:
	docker stop $(APP); docker rm $(APP)
	docker container prune --force
