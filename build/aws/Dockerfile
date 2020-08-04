FROM golang:1.14-alpine3.12 as build

ENV USERNAME=pogzyb
ENV CODE_REPO=${USERNAME}/website

COPY ../../website /go/src/${CODE_REPO}
WORKDIR /go/src/${CODE_REPO}

RUN apk add git && go get -v ./...

COPY ../../main.go /code/main.go
WORKDIR /code
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

# copy only the compiled program into final container
FROM alpine:3.12

RUN apk update && apk upgrade && apk add bash

ENV USERNAME=pogzyb
ENV CODE_REPO=${USERNAME}/website

COPY ../../website/templates /website/tmpl

COPY --from=build /code/main /website/main
WORKDIR /website
RUN chmod +x main

USER guest
CMD ["./main"]
