build:
	@go build -o ./bin/app

run: build
	@./bin/app

home:
	@curl -X GET "http://localhost:8080/"


