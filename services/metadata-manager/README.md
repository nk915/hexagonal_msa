
## before build
go mod tidy


## go build
cd services/metadata-manager

#### go excute
go run cmd/main.go

#### go build 
go build cmd/main.go


---


## docker build 

### docker image create
docker build -t metadata-manager .

### docker image run
docker run -P metadata-manager

or 

docker run -p 80:80 metadata-manager



--- 

## Testing

### go build http test
curl http://localhost:8080/services/kng

### docker http test
curl http://localhost:[docker_port or 80]/services/kng 
