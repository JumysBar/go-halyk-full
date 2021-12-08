FROM golang:1.16

WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/root/.cache go build -o /app/example -v ./cmd/main.go
ENTRYPOINT ["/app/example"]
