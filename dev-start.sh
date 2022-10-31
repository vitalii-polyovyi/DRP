#!/bin/bash
set -e
source ./dev-export-env.sh
set -x
export COMPOSE_PROJECT_NAME=drp
docker-compose -f docker-compose.local.yml up --build -d
