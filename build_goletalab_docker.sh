#!/bin/bash
set -e
IMAGE_NAME=${1:-units-task}
DOCKER_PLATFORM=${2:-linux/amd64}
DOCKER_BUILDKIT=1 docker build --platform "$DOCKER_PLATFORM" -f goletalab.Dockerfile -t "$IMAGE_NAME" .
echo "Docker image '$IMAGE_NAME' built successfully."
