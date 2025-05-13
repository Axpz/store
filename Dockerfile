# ---------- backend builder ----------
FROM golang:1.24-alpine AS backend-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o store ./cmd/

# ---------- final image using distroless ----------
FROM gcr.io/distroless/base-debian11

WORKDIR /app

ENV TZ=Asia/Shanghai

COPY --from=backend-builder /app/store .
COPY config.yaml .

EXPOSE 8080
ENTRYPOINT ["./store"]
