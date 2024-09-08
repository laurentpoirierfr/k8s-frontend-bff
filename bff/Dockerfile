FROM golang:1.23

RUN mkdir /src/
WORKDIR /src/
COPY . .
RUN go get ./...
RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM ubuntu:latest

COPY --from=0 /src/main /main
COPY --from=0 /src/configs /configs

EXPOSE 8080

CMD ["/main"]