# blockft-dex
Distributed Exchange

Currently runs off of the stellar quickstart container.

## Install
Install Go 1.14.2.
Install Stellar Quickstart Docker container.
```
docker pull stellar/quickstart
```

## Run
```
docker run --rm -it -p 8000:8000 -v --name stellar stellar/quickstart --standalone
go build test.go
./test.go
