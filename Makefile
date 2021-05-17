build:
	go build -o bin/main main.go

tidy:
	go mod tidy

test:
	go run main.go $(ARGS)

run:
	bin/main $(ARGS)