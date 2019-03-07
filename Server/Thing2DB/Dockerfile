#build golang on debian
FROM golang:1.12.0-stretch as builder

#download dep for golang
RUN go get github.com/golang/dep/cmd/dep

#copy Gopkg files and get dependencies
COPY ./Thing2DB/src/dal/Gopkg.lock ./Thing2DB/src/dal/Gopkg.toml /go/src/dal/
WORKDIR /go/src/dal/
RUN dep ensure -vendor-only

#copy Gopkg files and get dependencies
COPY ./Thing2DB/src/conn/Gopkg.lock ./Thing2DB/src/conn/Gopkg.toml /go/src/conn/
WORKDIR /go/src/conn/
RUN dep ensure -vendor-only

#copy Gopkg files and get dependencies
COPY ./Thing2DB/src/main/Gopkg.lock ./Thing2DB/src/main/Gopkg.toml /go/src/main/
WORKDIR /go/src/main/
RUN dep ensure -vendor-only

#copy source files to builder
COPY ./Thing2DB/src /go/src

#install and build from source
WORKDIR /go
RUN go install -v dal
RUN go install -v conn
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main main

#make final image
FROM debian:stretch as final
WORKDIR /bin

#copy program from builder and run
COPY --from=builder /go/bin/main main
COPY ./Thing2DB/config config
CMD ["chmod", "+x", "main"]
CMD [ "./main" ]
