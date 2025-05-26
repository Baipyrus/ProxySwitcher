#!/bin/bash

rm -rf ./build/*

sudo docker run --rm -it -v ./build:/app/build proxyswitcher:latest
# podman run --rm -it -v ./build:/app/build:z proxyswitcher:latest
