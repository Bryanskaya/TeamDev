all: swagger build
	
build: .
	mkdir -p logs
	mkdir -p temp
	go build -o .\temp .

swagger: ./controllers
	mkdir -p docs
	swag init

test:
	mockery --dir repository --all
	go test .\... -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html 

run: build
	.\temp\teamdev.exe
