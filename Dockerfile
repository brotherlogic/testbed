# syntax=docker/dockerfile:1
FROM golang:1.18 AS build

WORKDIR $GOPATH/src/github.com/brotherlogic/testbed

COPY go.mod ./
COPY go.sum ./

COPY *.go ./

RUN mkdir proto
COPY proto/*.go ./proto/

RUN go mod download

RUN CGO_ENABLED=0 go build -o /testbed

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /testbed /testbed

EXPOSE 8082

USER nonroot:nonroot

ENTRYPOINT ["/testbed"]