FROM golang:1.25.7-alpine AS build

WORKDIR /src

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /build/app ./cmd

FROM alpine:3.20

RUN adduser -D -u 10001 app

WORKDIR /app
COPY --from=build /build/app /app/main

EXPOSE 3333
USER app

ENTRYPOINT ["/app/main"]
