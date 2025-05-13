# ---------- frontend builder ----------
FROM node:23-alpine AS frontend-builder

WORKDIR /app

COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN npm install -g pnpm && pnpm install

COPY frontend/ ./
RUN pnpm build

# ---------- frontend runtime ----------
FROM node:23-alpine AS frontend-runtime

WORKDIR /app

ENV NODE_ENV=production
ENV TZ=Asia/Shanghai

COPY --from=frontend-builder /app/public ./public
COPY --from=frontend-builder /app/.next/static ./.next/static
COPY --from=frontend-builder /app/.next/standalone ./

EXPOSE 3000
CMD ["node", "server.js"]
