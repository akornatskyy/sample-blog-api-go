# Sample Blog API

[![Build Status](https://travis-ci.org/akornatskyy/sample-blog-api-go.svg?branch=master)](https://travis-ci.org/akornatskyy/sample-blog-api-go) [![Go Report Card](https://goreportcard.com/badge/github.com/akornatskyy/sample-blog-api-go)](https://goreportcard.com/report/github.com/akornatskyy/sample-blog-api-go)

A simple blog API written using go.

# Install

```sh
go get github.com/akornatskyy/sample-blog-api-go
```

# Run

```sh
cd $(go env GOPATH)/src/github.com/akornatskyy/sample-blog-api-go
go run main.go
```

# curl

```sh
curl -i -H 'Content-Type: application/json' \
  -d '{"username": "demo", "password": "password"}' \
  http://localhost:8080/signin

```

```
HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 21 Nov 2018 10:12:05 GMT
Content-Length: 20

{"username":"demo"}
```
