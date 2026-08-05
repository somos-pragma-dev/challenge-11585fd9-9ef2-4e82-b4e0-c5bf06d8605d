#!/bin/sh

set -e

RESPONSE=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:50051/health)

if [ $RESPONSE -ne 200 ]; then
  exit 1
fi

echo "Health check passed"