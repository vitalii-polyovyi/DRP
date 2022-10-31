#!/bin/bash
set -e
source ./dev-export-env.sh
set -x
docker-compose -f docker-compose.local.yml down
for dir in mongodb gridfs rabbitmq postgres
do
    unset $(grep -v '^#' $dir/.env | sed -E 's/(.*)=.*/\1/' | xargs)
done
