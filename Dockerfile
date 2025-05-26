FROM fedora:42 AS linux
# FROM ubuntu:plucky AS linux

WORKDIR /app
COPY . .

# RUN apt-get update && apt-get install -y \
#     libgl1-mesa-dev libxxf86vm-dev \
#     libxrandr-dev libxinerama-dev \
#     libx11-dev libxcursor-dev \
#     libxi-dev libglx-dev \
#     golang gcc

RUN echo "proxy=$http_proxy" | sudo tee -a /etc/dnf/dnf.conf
RUN sudo dnf update -y && sudo dnf install -y \
    libX11-devel libXcursor-devel libXrandr-devel \
    libXinerama-devel libXi-devel libGL-devel \
    libXxf86vm-devel mingw64-gcc-c++ golang

ENV GOBIN $GOROOT/bin
ENV GO_VERSION 1.24.2
ENV GO $GOBIN/go$GO_VERSION

RUN go install golang.org/dl/go$GO_VERSION@latest

RUN $GO download
RUN $GO mod download -x

ENV CGO_ENABLED 1
CMD $GO build -v -o build/


FROM linux AS windows

# TODO: RUN # install mingw-w64 tools
# RUN apt-get update && apt-get install -y mingw-w64

ENV GOOS windows
ENV GOARCH amd64
ENV CC "x86_64-w64-mingw32-gcc"
ENV CXX "x86_64-w64-mingw32-g++"

CMD $GO build -v -o build/
