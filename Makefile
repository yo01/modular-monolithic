run:
	@ printf "Starting Aplication... \n"
	@ cd '$(CURDIR)/src' &&  go run cmd/api/main.go 

tidy:
	cd '$(CURDIR)/src' && GOPRIVATE=git.motiolabs.com/library/motiolibs go get git.motiolabs.com/library/motiolibs
	cd '$(CURDIR)/src' && go mod tidy

test:
	cd '$(CURDIR)/src' && go test ./...

test-coverage:
	cd '$(CURDIR)/src' && go test ./... -coverprofile=coverage.out

build:
	@ printf "Building Aplication... \n"
	@ cd '$(CURDIR)/src' && go build \
		-trimpath  \
		-o engine \
		./app/
	@ echo "done"

go-generate: $(MOCKERY) ## Runs go generte ./...
	cd '$(CURDIR)/src' && go generate ./...