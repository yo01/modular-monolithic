#Builder
FROM golang:1.21.5-alpine3.18 AS build
WORKDIR /app
RUN apk update && apk add --no-cache gcc musl-dev git
COPY ./src/go.mod ./
COPY ./src/go.sum ./
RUN go mod download
RUN pwd
COPY ./src .
RUN go mod vendor
RUN go build -ldflags '-w -s' -a -o ./api ./cmd/api

# Distribution
FROM alpine:latest
RUN apk update
COPY --from=build /app/api /app/api
COPY --from=build /app/config.toml /app/config.toml
RUN  chmod +x /app/*
EXPOSE 3000
CMD ["/app/api"]
