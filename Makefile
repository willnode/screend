all: build

build:
	go build -o ./bin ./...
  
install:
	cd ./screend && go install
