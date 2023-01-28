# broker-service

```sh
docker build . -f broker-service.dockerfile -t broker
docker run --rm -it -p 8080:80 broker
```

```sh
curl -X POST http://localhost:8080/
```

### testing

```sh
go test -v ./cmd/api
```
