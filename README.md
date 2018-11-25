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

# Docker

```sh
docker run -it --rm -p 8080:8080 akorn/sample-blog-api-go
```

# curl

Validation error:


```sh
$ curl -i -H 'Content-Type: application/json' -d '{}' \
  http://localhost:8080/signin

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Thu, 22 Nov 2018 14:07:18 GMT
Content-Length: 265

{
  "errors": [{
      "domain": "signin",
      "type": "field",
      "location": "username",
      "reason": "required",
      "message": "Required field cannot be left blank."
    },
    {
      "domain": "signin",
      "type": "field",
      "location": "password",
      "reason": "required",
      "message": "Required field cannot be left blank."
    }
  ]
}
```

General error:

```sh
$ curl -i -H 'Content-Type: application/json' \
  -d '{"username":"js", "password": "password"}' \
  http://localhost:8080/signin

HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=UTF-8
Date: Thu, 22 Nov 2018 14:13:16 GMT
Content-Length: 167

{
  "errors": [{
    "domain": "signin",
    "type": "summary",
    "location": "user",
    "reason": "account locked",
    "message": "The account is locked. Contact system administrator, please."
  }]
}
```

Valid:

```sh
curl -i -H 'Content-Type: application/json' \
  -d '{"username": "demo", "password": "password"}' \
  http://localhost:8080/signin

HTTP/1.1 200 OK
Content-Type: application/json; charset=UTF-8
Date: Wed, 21 Nov 2018 10:12:05 GMT
Content-Length: 20

{"username":"demo"}
```
