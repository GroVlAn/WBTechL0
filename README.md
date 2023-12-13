# WB L0 Ordeer Service

## Run app on Unix, Linux, MacOs
```bash
make run
```

## Run on Windows

### Build apps
```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/orders ./cmd/orders/main.go;
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/dataGenerator ./cmd/dataGenerator/main.go;
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./build/client ./cmd/client/main.go;
```

### Docker compose build
```bash
docker-compose build
```

### Docker compose up
```bash
docker-compose --env-file ./.env  up -d
```
