FROM golang:1.17



RUN apt update -y \
    && apt-get install -y --no-install-recommends \
    build-essential gcc-mingw-w64 \
    ca-certificates \
    openssl \
    zip \
    clang

WORKDIR /src





