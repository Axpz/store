# ---------- frontend builder ----------
FROM node:20-alpine AS frontend-builder
WORKDIR /app/frontend

COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN npm install -g pnpm && pnpm install

COPY frontend/ ./
RUN pnpm build

# ---------- backend builder ----------
FROM golang:1.22-alpine AS backend-builder
WORKDIR /app/backend

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/store ./cmd/

# ---------- final image (distroless for security + performance) ----------
FROM gcr.io/distroless/base-debian11

WORKDIR /app

# timezone & certs
ENV TZ=Asia/Shanghai

# copy backend binary
COPY --from=backend-builder /app/store ./store
COPY config.yaml .

# copy frontend build (Next.js standalone mode 优先)
COPY --from=frontend-builder /app/frontend/public ./frontend/public
COPY --from=frontend-builder /app/frontend/.next/static ./frontend/.next/static
COPY --from=frontend-builder /app/frontend/.next/standalone ./frontend/

# copy start script
COPY docker-entrypoint.sh /usr/local/bin/
RUN chmod +x /usr/local/bin/docker-entrypoint.sh

# expose ports
EXPOSE 3000 8080

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]
