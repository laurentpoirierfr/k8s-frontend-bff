FROM golang:1.23

ENV GOPROXY="http://nexus-fabfonc.maif.local/repository/go-public,direct"
ENV GOSUMDB="off"

RUN mkdir /src/
WORKDIR /src/
COPY . .
RUN go get ./...
RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM ubuntu:latest

COPY --from=0 /src/main /main
COPY --from=0 /src/configs /configs

CMD ["/main"]