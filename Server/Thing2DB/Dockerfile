# Select latest golang container.
FROM golang:1.12.0-stretch as builder

# Install tools required to build the project
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers are only re-built when Gopkg files are updated
COPY ./Thing2DB/src/dal/Gopkg.lock ./Thing2DB/src/dal/Gopkg.toml /go/src/dal/
WORKDIR /go/src/dal/
# Install library dependencies
RUN dep ensure -vendor-only

COPY ./Thing2DB/src/conn/Gopkg.lock ./Thing2DB/src/conn/Gopkg.toml /go/src/conn/
WORKDIR /go/src/conn/
# Install library dependencies
RUN dep ensure -vendor-only

COPY ./Thing2DB/src/main/Gopkg.lock ./Thing2DB/src/main/Gopkg.toml /go/src/main/
WORKDIR /go/src/main/
# Install library dependencies
RUN dep ensure -vendor-only


# Copy all project and build it
# This layer is rebuilt when ever a file has changed in the project directory
COPY ./Thing2DB/src /go/src

# Build all the projects
WORKDIR /go
RUN go install -v dal
RUN go install -v conn
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main main

# This results in a single layer image
FROM debian:stretch as final
WORKDIR /bin

# Entry point is the Go webserver
COPY --from=builder /go/bin/main main
COPY ./Thing2DB/config config
CMD ["chmod", "+x", "main"]
CMD [ "./main" ]
