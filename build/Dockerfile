FROM golang:1.23 AS builder

WORKDIR /shortener

COPY go.mod go.sum ./

RUN go mod download -x

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" /shortener/cmd/shortener/shortener.go

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

COPY --from=builder /shortener/shortener ./

EXPOSE  2821

ENTRYPOINT ["/app/shortener"]
