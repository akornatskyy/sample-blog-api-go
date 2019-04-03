FROM golang:alpine as b

ENV APPDIR=/go/src/github.com/akornatskyy/sample-blog-api-go

ADD . $APPDIR
WORKDIR $APPDIR

RUN set -ex \
    \
    && apk add --no-cache git upx \
    \
    && go get -v -d \
    && CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' \
      -o /go/bin/sample-blog-api \
    && upx -q --brute /go/bin/sample-blog-api

FROM scratch

LABEL maintainer="Andriy Kornatskyy <andriy.kornatskyy@live.com>"

ENV KEY=

COPY --from=b /go/bin/sample-blog-api /app/
ADD samples.json /app/
WORKDIR /app

EXPOSE 8080

CMD ["./sample-blog-api"]
