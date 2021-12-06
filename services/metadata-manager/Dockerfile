# syntax=docker/dockerfile:1
FROM golang:1.16-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app

## Add this go mod download command to pull in any dependencies
## Our project will now successfully build with the necessary go libraries included.
RUN go mod download

## go runtime env setting
ENV APPENV production 

## our newly created binary executable
RUN go build -o metadata-manager ./cmd/

EXPOSE 80
CMD [ "/app/metadata-manager" ]