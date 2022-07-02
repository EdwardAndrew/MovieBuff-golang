##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY ./pkg ./pkg


RUN go build -o /moviebuff

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /
COPY --from=build /moviebuff /moviebuff
COPY .env ./
USER nonroot:nonroot

ENTRYPOINT ["/moviebuff"]