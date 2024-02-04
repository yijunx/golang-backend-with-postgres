FROM ubuntu:latest

ENV GO_VERSION=1.21.5
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=arm64

# well below has to be in one line!
RUN apt-get update && apt-get install -y wget git gcc make 

RUN wget -P /tmp "https://dl.google.com/go/go${GO_VERSION}.linux-${GOARCH}.tar.gz"
RUN tar -C /usr/local -xzf "/tmp/go${GO_VERSION}.linux-${GOARCH}.tar.gz"
RUN rm "/tmp/go${GO_VERSION}.linux-${GOARCH}.tar.gz"

# ENV GOPATH /opt/go/code
ENV PATH /usr/local/go/bin:$PATH
# # here we add the code folder into path, then we dont need go.mod anymore
# RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

WORKDIR /opt/go/code
ENTRYPOINT [ "/bin/sh", "-c" ]