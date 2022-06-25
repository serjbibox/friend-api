BINARY_NAME=friend-api

build:
	docker build -t friend-api:latest -f Dockerfile .
run:	
	docker run -d -p 8080:8080 -p 5432:5432 friend-api
