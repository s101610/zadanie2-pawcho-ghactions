# syntax=docker/dockerfile:1.7

# Pierwszy etap kompiluje prostą aplikację Go.
# Wartości TARGETOS oraz TARGETARCH są ustawiane automatycznie przez buildx.
FROM --platform=$BUILDPLATFORM golang:1.25.10-alpine3.23 AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY src ./src
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH \
    go build -trimpath -ldflags="-s -w" -o /dist/server ./src

# Drugi etap tworzy mały obraz wynikowy.
# Obraz scratch ogranicza liczbę zbędnych plików w finalnym kontenerze.
FROM scratch

COPY --from=builder /dist/server /server
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/server"]
