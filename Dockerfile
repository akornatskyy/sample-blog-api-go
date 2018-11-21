FROM golang:alpine as b

ENV APPDIR=/go/src/github.com/akornatskyy/sample-blog-api-go

ADD . $APPDIR
WORKDIR $APPDIR

RUN set -ex \
    \
    && CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' \
      -o /go/bin/sample-blog-api

FROM scratch

LABEL maintainer="Andriy Kornatskyy <andriy.kornatskyy@live.com>"

COPY --from=b /go/bin/sample-blog-api /app/
ADD user-samples.json /app
WORKDIR /app

EXPOSE 8080

CMD ["./sample-blog-api"]
