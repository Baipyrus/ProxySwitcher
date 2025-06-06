FROM ubuntu:24.04 AS linux

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y \
    libgl1-mesa-dev libxxf86vm-dev \
    libxrandr-dev libxinerama-dev \
    libx11-dev libxcursor-dev \
    libxi-dev libglx-dev \
    golang gcc

ENV GOBIN $GOROOT/bin
ENV GO_VERSION 1.23.3
ENV GO $GOBIN/go$GO_VERSION

RUN go install golang.org/dl/go$GO_VERSION@latest

RUN $GO download
RUN $GO mod download -x

ENV CGO_ENABLED 1
CMD $GO build -v -o build/


FROM linux AS windows

RUN apt-get update && apt-get install -y mingw-w64

ENV GOOS windows
ENV GOARCH amd64
ENV CC "x86_64-w64-mingw32-gcc"
ENV CXX "x86_64-w64-mingw32-g++"

CMD $GO build -v -o build/
