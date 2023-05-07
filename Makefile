all: build

build:
	cd cmd/screend \
		&& go generate \
		&& go build -o ../../bin/screend \
  
install:
	cd cmd/screend \
		&& go generate \
		&& go install \
