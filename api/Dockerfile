FROM quay.io/goswagger/swagger:v0.24.0 AS generate

ENV GOPATH /go
COPY . /go/src/github.com/ndebuhr/billing-service/api
WORKDIR /go/src/github.com/ndebuhr/billing-service/api
RUN go mod init github.com/ndebuhr/billing-service && \
    swagger generate server -f swagger.yml && \
    go get -v -d ./...

FROM golang:1.13 as build

ENV GOPATH /go
RUN rm -rf /go
COPY --from=generate /go /go
WORKDIR /go/src/github.com/ndebuhr/billing-service/api/cmd/billing-portal-server
RUN go build -o /opt/app .

FROM gcr.io/distroless/base

COPY --from=build /opt/app /opt/app

EXPOSE 80

ENTRYPOINT ["/opt/app"]
CMD ["--host","0.0.0.0","--port","80"]