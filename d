#!/bin/sh

docker compose -f docker/docker-compose.yaml exec app "$@"
