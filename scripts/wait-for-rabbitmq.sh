#!/bin/bash

set -e

host="$1"
shift
cmd="$@"

until rabbitmqctl status; do
  >&2 echo "RabbitMQ is unavailable - sleeping"
  sleep 1
done

>&2 echo "RabbitMQ is up - executing command"
exec $cmd 