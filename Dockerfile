FROM golang:alpine as builder

WORKDIR /GoComment-service

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/

FROM alpine

WORKDIR /GoComment-service

COPY --from=builder /GoComment-service/main /GoComment-service/main
COPY --from=builder /GoComment-service/pkg/config/envs/*.env /GoComment-service/

RUN chmod +x /GoComment-service/main

CMD ["./main"]