FROM golang:1.12.1-alpine AS builder

RUN apk add --no-cache ca-certificates git

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN cd /src/cmd/reflektor && \
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /bin/reflektor

FROM alpine:latest

WORKDIR /bin
COPY --from=builder /bin/reflektor /bin/reflektor

ENTRYPOINT ["/bin/reflektor"]