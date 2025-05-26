#!/bin/bash

sudo docker build \
    --build-arg HTTP_PROXY=$http_proxy \
    --build-arg HTTPS_PROXY=$http_proxy \
    --build-arg http_proxy=$http_proxy \
    --build-arg https_proxy=$http_proxy \
    -t proxyswitcher:latest \
    -f Dockerfile . \
    $@
# podman build -t proxyswitcher:latest -f Dockerfile . $@
