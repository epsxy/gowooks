build:
	go build -o ./bin/gowooks ./src/main.go

run:
	./bin/gowooks

test:
	go test ./...

dockerbuild:
	docker build --tag gowooks:latest .

dockerun:
	docker run -p 8080:80 -it gowooks:latest
