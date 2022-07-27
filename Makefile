SERVICE_NAME=example

.PHONY: docker up up-fg down clean

${SERVICE_NAME}_darwin: clean
	@echo "Making ${SERVICE_NAME}_darwin..."
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ./${SERVICE_NAME} -ldflags "$(LDFLAGS)" ./main.go
	go mod tidy

${SERVICE_NAME}_linux: clean
	@echo "Making ${SERVICE_NAME}_linux..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./${SERVICE_NAME} -ldflags "$(LDFLAGS)" ./main.go
	strip ${SERVICE_NAME}
	go mod tidy

docker: ${SERVICE_NAME}_linux
	docker build -t ${SERVICE_NAME} .

up:
	docker-compose up -d

up-fg:
	docker-compose up

down:
	docker-compose down

clean:
	rm -f ./${SERVICE_NAME}