FROM golang:1.14-alpine3.12

ENV USERNAME=pogzyb
ENV CODE_REPO=${USERNAME}/website

COPY website /go/src/${CODE_REPO}
WORKDIR /go/src/${CODE_REPO}

RUN apk add git && go get -v ./...
RUN go get github.com/pilu/fresh

COPY website/templates /website/templates
COPY runner.conf /website/runner.conf
COPY main.go /website/main.go

WORKDIR /website

CMD ["fresh", "-c", "runner.conf"]
