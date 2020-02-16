FROM golang:1.13-alpine as builder

# Force Go to use the cgo based DNS resolver. This is required to ensure DNS
# queries required to connect to linked containers succeed.
ENV GODEBUG netdns=cgo

# Install dependencies and install/build lnd.
RUN apk add --no-cache --update alpine-sdk \
    git \
    make 

# Copy in the local repository to build from.
COPY . /go/src/github.com/YusukeShimizu/Candy_Scrape

RUN cd /go/src/github.com/YusukeShimizu/Candy_Scrape \
    &&  pwd \
    &&  go build main.go

# Expose lnd ports (server, rpc).
EXPOSE 3000

# Add bash.
RUN apk add --no-cache \
    bash

WORKDIR /go/src/github.com/YusukeShimizu/Candy_Scrape