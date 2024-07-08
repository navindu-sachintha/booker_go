# BUILD Stage
FROM golang:1.22-alpine as builder

WORKDIR /project/bookstoreapi

COPY go.* ./

RUN go mod download

COPY . .
RUN go mod tidy && go mod vendor
WORKDIR /project/bookstoreapi/cmd/main
RUN go build -o ../../bookstoreapp

# Execute Stage
FROM alpine:latest

COPY --from=builder /project/bookstoreapi/bookstoreapp /project/bookstoreapi/bookstoreapp

EXPOSE 8080
ENTRYPOINT [ "/project/bookstoreapi/bookstoreapp" ]