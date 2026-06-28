FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -ldflags "-X github.com/kazyamaz200/agentos/internal/cli.Version=${VERSION}" -o agentos ./cmd/agentos

FROM alpine:3.19

RUN apk add --no-cache ca-certificates tzdata bash git

WORKDIR /workspace
COPY --from=builder /build/agentos /usr/local/bin/agentos

EXPOSE 8080

ENTRYPOINT ["agentos"]
CMD ["--help"]
