FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o store-manager ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/store-manager .
EXPOSE 8080
ENTRYPOINT ["./store-manager"]