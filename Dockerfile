FROM golang:1.13-alpine

RUN apk add --no-cache make ncurses

COPY . $GOPATH/src/github.com/wtfutil/wtf

ENV GOPROXY=https://gocenter.io
ENV GO111MODULE=on
ENV GOSUMDB=off

WORKDIR $GOPATH/src/github.com/wtfutil/wtf

ENV PATH=$PATH:./bin

RUN make build

ENTRYPOINT "wtfutil" 
