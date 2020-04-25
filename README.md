# blockft-dex
Distributed Exchange

Currently runs off of the stellar quickstart container.

## Install
Install Go 1.14.2.
Install Stellar Quickstart Docker container.
```
docker pull stellar/quickstart
git clone https://github.com/thegajan/blockft-dex.git
cd blockft-dex/src/app
go install
```

## Run
```
docker run --rm -it -p "8000:8000" -v "/home/user/stellar:/opt/stellar" --name stellar stellar/quickstart --standalone
app
```
