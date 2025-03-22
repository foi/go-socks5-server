#!/usr/bin/env bash

export VERSION=$(cat VERSION)

export CGO_ENABLED=0

GOOS=linux GOARCH=amd64 go build -o build/go-socks5-server-amd64 -ldflags="-s -w -X main.Version=v$VERSION" main.go
GOOS=linux GOARCH=arm64 go build -o build/go-socks5-server-arm64 -ldflags="-s -w -X main.Version=v$VERSION" main.go

mkdir -p build/tmp/usr/bin
mkdir -p build/tmp/etc/systemd/system
mkdir -p build/tmp/usr/lib/sysusers.d

cp go-socks5-server.service build/tmp/etc/systemd/system
cp sysusers.conf build/tmp/usr/lib/sysusers.d/go-socks5-server.conf
cp go-socks5-server.config.json.example build/tmp/etc/go-socks5-server.config.json

cd build

cp go-socks5-server-amd64 tmp/usr/bin/go-socks5-server

cd ..

docker run --rm -it -e VERSION=$VERSION -v $PWD:/app -w /app/build foifirst/fpm:ruby3.3-fedora41 fpm -s dir -t rpm -a x86_64 -C /app/build/tmp --description "go-socks5-server" --name go-socks5-server --post-install /app/post-install.sh --version $VERSION --iteration 1 .

docker run --rm -it -e VERSION=$VERSION -v $PWD:/app -w /app/build foifirst/fpm:ruby3.3-fedora41 fpm -s dir -t deb -a x86_64 -C /app/build/tmp --description "go-socks5-server" --name go-socks5-server --post-install /app/post-install.sh --version $VERSION --iteration 1 .

cd build

cp go-socks5-server-arm64 tmp/usr/bin/go-socks5-server

cd ..

docker run --rm -it -e VERSION=$VERSION -v $PWD:/app -w /app foifirst/fpm:ruby3.3-fedora41 fpm -s dir -t deb -a arm64 -C /app/build/tmp --description "go-socks5-server" --name go-socks5-server --post-install /app/post-install.sh --version $VERSION --iteration 1 .

docker run --rm -it -e VERSION=$VERSION -v $PWD:/app -w /app foifirst/fpm:ruby3.3-fedora41 fpm -s dir -t rpm -a arm64 -C /app/build/tmp --description "go-socks5-server" --name go-socks5-server --post-install /app/post-install.sh --version $VERSION --iteration 1 .
