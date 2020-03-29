# go-book
## Run with go
1. Install [Go](https://golang.org/doc/install).
2. Install [MongoDB](https://docs.mongodb.com/manual/installation/).
3. Run MongoDB:
4. Run app:
```
  go run main.go
```
5. [http://localhost:1323](http://localhost:1323)
## Run with docker compose
1. Install [Docker](https://docs.docker.com/get-docker/).
2. Run app:
```
  docker-compose up -d --build
```
3. [http://localhost:1323](http://localhost:1323)
## Run test
```
  go test ./ ./test -v -cover
```