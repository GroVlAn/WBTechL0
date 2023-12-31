.PHONY:

#Build producer application
g-build-producer:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/orders ./cmd/orders/main.go;

#Build subscriber application
g-build-subscriber:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/dataGenerator ./cmd/dataGenerator/main.go;

#Build client application
g-build-client:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/client ./cmd/client/main.go;

#Docker compose build
compose-build:
	docker-compose build

#Docker compose run
compose-run:
	docker-compose --env-file ./.env  up -d

#Run applications
run: g-build-producer g-build-subscriber g-build-client compose-build compose-run
