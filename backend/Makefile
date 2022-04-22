all: swagger build test

build: . logs temp
	go build -o .\temp .

swagger: ./controllers docs
	swag init

mockery:
	mockery --dir repository --all

test: mockery logs
	go test -v -coverprofile=logs/c.out -coverpkg ./models/ ./...
	go tool cover -html=logs/c.out -o logs/coverage.html 

run: build
	.\temp\teamdev.exe

# Creating directories
DIRS?=logs temp
$(DIRS):
	mkdir -p $@
