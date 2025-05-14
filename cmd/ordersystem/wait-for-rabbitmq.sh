#!/bin/sh
# wait-for-rabbitmq.sh

set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until nc -z "$host" "$port"; do
  echo "Esperando RabbitMQ em $host:$port..."
  sleep 1
done

echo "RabbitMQ está pronto!"
exec $cmd 