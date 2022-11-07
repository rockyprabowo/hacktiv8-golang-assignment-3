# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -tags prod -o /h8-assignment-3

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /h8-assignment-2 /h8-assignment-3

EXPOSE 5001

USER nonroot:nonroot

ENTRYPOINT ["/h8-assignment-3"]
