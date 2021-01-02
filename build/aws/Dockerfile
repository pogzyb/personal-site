FROM golang:1.14-alpine3.12 as build

ENV USERNAME=pogzyb
ENV CODE_REPO=${USERNAME}/website

COPY ./pkg/website /go/src/${CODE_REPO}
WORKDIR /go/src/${CODE_REPO}

RUN apk add \
    git \
    build-base \
    && go get gorm.io/gorm \
    && go get -v ./...

COPY ./cmd/main.go /app/main.go
WORKDIR /app
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o main main.go

# copy only the compiled program into final container
FROM alpine:3.12

RUN apk update && apk upgrade && apk add bash

ENV USERNAME=pogzyb
ENV CODE_REPO=${USERNAME}/website

COPY ./web /app/web
COPY --from=build /app/main /app/main
WORKDIR /app
RUN chmod +x main
USER guest
CMD ["./main"]
