FROM ubuntu:latest AS linux

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y \
    libgl1-mesa-dev libxxf86vm-dev \
    libxrandr-dev libxinerama-dev \
    libx11-dev libxcursor-dev \
    libxi-dev libglx-dev \
    golang gcc

ENV GOBIN $GOROOT/bin
ENV GO_VERSION 1.24.2
ENV GO $GOBIN/go$GO_VERSION

RUN go install golang.org/dl/go$GO_VERSION@latest

RUN $GO download
RUN $GO mod download -x

ENV CGO_ENABLED 1
CMD $GO build -v -o build/
