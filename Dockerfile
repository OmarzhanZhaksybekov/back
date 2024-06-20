FROM golang:1.21.1-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

COPY ["go.mod", "./" ]
COPY ["go.sum", "./"]
RUN go mod download

COPY  . ./
RUN go build -o ./bin/app cmd/main.go
CMD [ "./bin/app" ]