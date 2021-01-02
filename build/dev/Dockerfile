FROM golang:1.14-alpine3.12

ENV USERNAME=pogzyb
ENV CODE_REPO=${USERNAME}/website

COPY ./pkg/website /go/src/${CODE_REPO}
WORKDIR /go/src/${CODE_REPO}

RUN apk add \
    git \
    build-base \
    && go get gorm.io/gorm \
    && go get -v ./... \
    && go get github.com/pilu/fresh

COPY ./web /app/web
COPY runner.conf /app/runner.conf
COPY ./cmd/main.go /app/main.go

WORKDIR /app

CMD ["fresh", "-c", "runner.conf"]
