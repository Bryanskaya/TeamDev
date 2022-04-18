all: swagger build
	
build: .
	mkdir -p logs
	mkdir -p temp
	go build -o .\temp .

swagger: ./controllers
	mkdir -p docs
	swag init

run: build
	.\temp\teamdev.exe
