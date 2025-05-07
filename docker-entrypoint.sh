#!/bin/sh

set -e

MODE="$1"

if [ "$MODE" = "frontend" ]; then
  echo "start frontend..."
  cd /app/frontend
  pnpm start
elif [ "$MODE" = "backend" ]; then
  echo "start backend..."
  ./store
else
  echo "Usage: $0 [frontend|backend]"
  exit 1
fi